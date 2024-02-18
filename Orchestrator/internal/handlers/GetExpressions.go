package handlers

import (
	"context"
	"net/http"

	"github.com/Rehart-Kcalb/DAEC/Orchestrator/internal/db"
	"github.com/Rehart-Kcalb/DAEC/Orchestrator/internal/types"
)

type GetExpressions struct {
	db.Queries
}

func (g GetExpressions) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	expressions, err := g.Queries.GetExpressions(context.TODO())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	types.NewSuccessResponse(http.StatusOK, struct {
		Expressions any `json:"expressions"`
	}{expressions}).Respond(w)
}
