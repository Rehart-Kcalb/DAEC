package handlers

import (
	"encoding/json"
	"net/http"
)

type AddExpression struct{}

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
		ID         string `json:"id"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call parser from ../parser/parser.go
	err = parser.processExpression(data.Expression) // Assuming parseExpression is exported from ../parser/parser.go
	if err != nil {
		// Handle parsing errors
		http.Error(w, "Error parsing expression", http.StatusBadRequest)
		return
	}

	// Indicate success
	w.WriteHeader(http.StatusOK)
}
