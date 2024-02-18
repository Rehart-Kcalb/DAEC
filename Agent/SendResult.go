package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func sendResult(result Result) {
	// Simulate sending result to server (replace with actual HTTP request)
	url := orchestratorURL + "/api/task"
	payload, _ := json.Marshal(result)
	_, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("Error sending result: %v\n", err)
	}
}
