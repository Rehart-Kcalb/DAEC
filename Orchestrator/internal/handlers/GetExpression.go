package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/Rehart-Kcalb/DAEC/Orchestrator/internal/db"
	"github.com/Rehart-Kcalb/DAEC/Orchestrator/internal/types"
)

type GetExpression struct {
	*db.Queries
}

func (g GetExpression) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("Not valid id"))
		if err != nil {
			log.Println("Error: while writing data to response " + err.Error())
			return
		}
	}

	expression, err := g.Queries.GetExpression(context.TODO(), int64(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	types.NewSuccessResponse(http.StatusOK, struct {
		Expression any `json:"expression"`
	}{expression}).Respond(w)
}
