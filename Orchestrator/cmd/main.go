package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/Rehart-Kcalb/DAEC/Orchestrator/internal/db"
	"github.com/Rehart-Kcalb/DAEC/Orchestrator/internal/handlers"
	"github.com/Rehart-Kcalb/DAEC/Orchestrator/internal/types"
	"github.com/jackc/pgx/v5/pgxpool"
)

var agents types.AgentsMutex = types.AgentsMutex{Agents: make(map[string]time.Time), Mutex: &sync.Mutex{}}
var deadTime float64 = 10

func main() {

	/*if err := godotenv.Load(); err != nil {*/
	/*log.Fatal("Error loading .env ", err)*/
	/*}*/

	db_url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("DB_NAME"))

	conn, err := pgxpool.New(context.Background(), db_url)
	if err != nil {
		log.Println(db_url)
		log.Fatal("Fault connect to database:", err.Error())
	}

	queries := db.New(conn)

	Agents, err := queries.GetAgents(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("App started normally")

	for _, agent := range Agents {
		agents.Agents[agent] = time.Now()
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/expression", handlers.AddExpression{Queries: queries}.ServeHTTP)
	mux.HandleFunc("GET /api/expressions", handlers.GetExpressions{Queries: *queries}.ServeHTTP)
	mux.HandleFunc("GET /api/expressions/{id}", handlers.GetExpression{Queries: queries}.ServeHTTP)
	mux.HandleFunc("GET /api/operations/", handlers.GetOperations{Queries: *queries}.ServeHTTP)
	// TODO: Change operations route
	mux.HandleFunc("POST /api/operation", handlers.UpdateOperation{Queries: queries}.ServeHTTP)
	mux.HandleFunc("GET /api/monitor", handlers.Monitor{Queries: queries}.ServeHTTP)
	mux.HandleFunc("GET /api/task", handlers.GiveTask{Queries: queries, Agents: agents}.ServeHTTP)
	mux.HandleFunc("POST /api/task", handlers.GetResult{Queries: queries}.ServeHTTP)
	mux.HandleFunc("POST /api/ping", (&handlers.ProcessPing{Agents: agents}).ServeHTTP)

	go func() {
		for {
			agents.Mutex.Lock()
			for k, l := range agents.Agents {
				if time.Since(l).Seconds() < deadTime {
					continue
				}
				expressions, err := queries.GetAgentExpressions(context.Background(), k)
				if err != nil {
					log.Println("Error, while getting tasks for freeing", err)
					continue
				}
				if err := queries.ClearTakenForAgent(context.Background(), k); err != nil {
					log.Println("Error, while freeing tasks", err)
					continue
				}
				for _, expr := range expressions {
					if err := queries.UpdateExpressionStatus(context.Background(), db.UpdateExpressionStatusParams{ID: expr.ExpressionID, Status: db.ExpressionStatusWait}); err != nil {
						log.Println("Error, while changing tasks for freed tasks", err)
						continue
					}
				}
				delete(agents.Agents, k)
			}
			agents.Mutex.Unlock()
		}
	}()

	http.ListenAndServe(":8080", mux)
}
