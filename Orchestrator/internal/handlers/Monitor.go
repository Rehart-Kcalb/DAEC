package handlers

import (
	"context"
	"net/http"

	"github.com/Rehart-Kcalb/DAEC/Orchestrator/internal/db"
	"github.com/Rehart-Kcalb/DAEC/Orchestrator/internal/types"
)

type Monitor struct {
	*db.Queries
}

func (m Monitor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	agents, err := m.Queries.Monitor(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	types.NewSuccessResponse(http.StatusOK, struct {
		Monitor any `json:"monitor"`
	}{agents})
}
