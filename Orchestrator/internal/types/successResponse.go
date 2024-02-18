package types

import (
	"encoding/json"
	"log"
	"net/http"
)

type successResponse struct {
	Status int `json:"status"`
	Data   any `json:"data"`
}

func NewSuccessResponse(status int, data any) *successResponse {
	return &successResponse{Status: status, Data: data}
}

func (s *successResponse) Respond(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s.Status)
	err := json.NewEncoder(w).Encode(s)
	if err != nil {
		log.Printf("Error, while writing response, err: %s\n", err.Error())
	}
}
