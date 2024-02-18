package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Rehart-Kcalb/DAEC/Orchestrator/internal/types"
)

type ProcessPing struct {
	Agents types.AgentsMutex
}

func (p *ProcessPing) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.Agents.Mutex.Lock()
	defer p.Agents.Mutex.Unlock()
	var data struct {
		Agent string `json:"agent"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	log.Println(data)
	p.Agents.Agents[data.Agent] = time.Now()
}
