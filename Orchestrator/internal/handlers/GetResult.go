package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Rehart-Kcalb/DAEC/Orchestrator/internal/db"
)

type GetResult struct {
	*db.Queries
}

func (g GetResult) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var data struct {
		ExpressionId    int64   `json:"expression_id"`
		SubExpressionId int64   `json:"sub_expression_id"`
		Result          float64 `json:"result"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	log.Println(data)

	if data.SubExpressionId == -1 {
		err := g.Queries.InputResultExpression(context.Background(), db.InputResultExpressionParams{ID: data.ExpressionId, Result: data.Result})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = g.Queries.DeleteTaken(context.Background(), data.ExpressionId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	err = g.Queries.InputResultSubExpression(context.Background(), db.InputResultSubExpressionParams{ID: data.SubExpressionId, Result: data.Result})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
