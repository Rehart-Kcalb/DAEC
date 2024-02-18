package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Rehart-Kcalb/DAEC/Orchestrator/internal/db"
	"github.com/Rehart-Kcalb/DAEC/Orchestrator/internal/handlers"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env")
	}

	db_url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("DB_NAME"))

	conn, err := pgx.Connect(context.Background(), db_url)
	if err != nil {
		log.Println(db_url)
		log.Fatal("Fault connect to database:", err.Error())
	}

	queries := db.New(conn)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/expression", handlers.AddExpression{}.ServeHTTP)
	mux.HandleFunc("GET /api/expressions", handlers.GetExpressions{Queries: *queries}.ServeHTTP)
	mux.HandleFunc("GET /api/expressions/{id}", handlers.GetExpression{Queries: queries}.ServeHTTP)
	mux.HandleFunc("GET /api/operations/", handlers.GetOperations{Queries: *queries}.ServeHTTP)
	mux.HandleFunc("GET /api/task", handlers.GiveTask)
	mux.HandleFunc("POST /api/task", handlers.GetResult)
	mux.HandleFunc("POST /api/ping", handlers.ProcessPing)

	http.ListenAndServe(":8080", mux)
}
