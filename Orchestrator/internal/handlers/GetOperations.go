package handlers

import (
	"context"
	"net/http"

	"github.com/Rehart-Kcalb/DAEC/Orchestrator/internal/db"
	"github.com/Rehart-Kcalb/DAEC/Orchestrator/internal/types"
)

type GetOperations struct {
	db.Queries
}

func (g GetOperations) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	operations, err := g.Queries.GetOperations(context.TODO())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	types.NewSuccessResponse(http.StatusOK, struct {
		Operations any `json:"operations"`
	}{operations}).Respond(w)
}
