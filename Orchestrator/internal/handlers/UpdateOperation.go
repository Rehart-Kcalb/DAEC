package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Rehart-Kcalb/DAEC/Orchestrator/internal/db"
)

type UpdateOperation struct {
	*db.Queries
}

func (u UpdateOperation) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var data struct {
		Operations []db.Operation `json:"operations"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	for _, val := range data.Operations {
		err := u.Queries.UpdateOperation(context.Background(), db.UpdateOperationParams{Cost: val.Cost, Operation: val.Operation})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}
