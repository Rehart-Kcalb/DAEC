package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Rehart-Kcalb/DAEC/Orchestrator/internal/db"
	"github.com/Rehart-Kcalb/DAEC/Orchestrator/internal/parser"
	"github.com/jackc/pgx/v5/pgtype"
)

type AddExpression struct {
	*db.Queries
}

/*
	AddExpression handler handle /api/expression path.
	I need it get expression from POST body and also unique id for expression.
	Check it with ../parser/parser.go, if no error return http.StatusOK.
	Else if error with parsing return 400.
	Else if error not related to parsing return 500
*/

func (s AddExpression) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Decode expression and ID from request body
	var data struct {
		Expression string `json:"expression"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	operations, err := s.Queries.GetOperations(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var op_map map[string]int = make(map[string]int)
	for _, op := range operations {
		op_map[string(op.Operation)] = int(op.Cost)
	}

	// Call parser from ../parser/parser.go
	// TODO: Make adding subexprressions
	now := &pgtype.Timestamptz{}
	err = now.Scan(time.Now())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	id, err := s.Queries.InsertExpression(context.Background(), db.InsertExpressionParams{Expression: data.Expression, Status: "wait", CreatedAt: *now})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	subexpressions, err := parser.ProcessExpression(data.Expression, op_map) // Assuming parseExpression is exported from ../parser/parser.go
	if err != nil {
		// Handle parsing errors
		s.Queries.UpdateExpressionStatus(context.Background(), db.UpdateExpressionStatusParams{ID: id, Status: db.ExpressionStatusInvalid})
		http.Error(w, fmt.Sprintf("Error parsing expression. \nId:%d", id), http.StatusBadRequest)
		return
	}

	counter := 1
	for !subexpressions.IsEmpty() {
		expression, _ := subexpressions.Pop()
		log.Println(expression)
		err = s.Queries.InsertSubExpression(context.Background(), db.InsertSubExpressionParams{Operand1: expression.Operand1, Operand2: expression.Operand2, Operation: db.OperationSymbol(expression.Operator), Cost: int32(expression.Cost), ExpressionID: id, ExecOrder: int32(counter)})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		counter++
	}

	// Indicate success
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%d\n", id)
}
