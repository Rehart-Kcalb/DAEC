package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Rehart-Kcalb/DAEC/Orchestrator/internal/db"
	"github.com/Rehart-Kcalb/DAEC/Orchestrator/internal/types"
)

type GiveTask struct {
	*db.Queries
	Agents types.AgentsMutex
}

func (g GiveTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.Agents.Lock()
	defer g.Agents.Unlock()
	agent := r.Header.Get("X-AGENT-NAME")
	if agent == "" {
		log.Println("Empty AGENT name")
	}
	calculator := r.FormValue("calculator")
	Calculator, err := strconv.Atoi(calculator)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Calculator value is not valid int"))
		return
	}

	expression, err := g.Queries.GetTask(context.TODO())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	subExpressions, err := g.Queries.GetSubexpressions(context.TODO(), expression.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = g.Queries.InsertTaken(context.TODO(), db.InsertTakenParams{Agent: agent, Calculator: int32(Calculator), ExpressionID: expression.ID})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = g.Queries.UpdateExpressionStatus(context.Background(), db.UpdateExpressionStatusParams{ID: expression.ID, Status: db.ExpressionStatusProcessing})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	g.Agents.Agents[agent] = time.Now()
	types.NewSuccessResponse(http.StatusOK, struct {
		Expression    any `json:"expression"`
		SubExpression any `json:"sub_expressions"`
	}{expression, subExpressions}).Respond(w)
}
