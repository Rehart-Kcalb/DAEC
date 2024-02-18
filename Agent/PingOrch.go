package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func pingOrchestrator(agent_name string) {
	// Simulate pinging orchestrator with processed expression (replace with actual HTTP request)
	url := orchestratorURL + "/api/ping"
	// payload, _ := json.Marshal(expressions)
	payload, _ := json.Marshal(struct {
		Agent string `json:"agent"`
	}{agent_name})
	_, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("Error pinging orchestrator: %v\n", err)
	}
}
