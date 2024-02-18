package main

import (
	"io"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

var expressions []Expression

var GoroutineCounter int

var orchestratorURL string

func main() {

	orchestratorURL = os.Getenv("orchURL")

	if orchestratorURL == "" {
		log.Println("There is no orchestratorURL environment variable")
		return
	}

	maxGo, err := strconv.Atoi(os.Getenv("MAX_GOROUTINE"))
	if err != nil {
		log.Fatal("Problem with env variable MAX_GOROUTINE")
	}
	maxGoroutines := maxGo
	GoroutineCounter := make(chan int, maxGoroutines)

	var wg sync.WaitGroup

	// Orchestrator endpoint to receive expressions
	GetTaskEndPoint := orchestratorURL + "/api/task"

	// Goroutine to fetch expressions from orchestrator
	go func() {
		for {
			// Retrieve expression from orchestrator with a timeout
			if len(GoroutineCounter) >= maxGoroutines {
				time.Sleep(1 * time.Second)
				continue
			}
			exps, err := getExpressionFromOrchestrator(GetTaskEndPoint, len(GoroutineCounter)+1)
			if err != nil {
				if err == io.EOF {
					time.Sleep(1 * time.Second)
					continue
				}
				log.Printf("Error getting expression from orchestrator: %v", err)
				continue
			}
			GoroutineCounter <- 1 // Increment GoroutineCounter
			wg.Add(1)

			// Process expression in a separate goroutine
			go func(exp []Expression) {
				defer func() {
					<-GoroutineCounter // Decrement GoroutineCounter
					wg.Done()
				}()
				processExpression(exp)
			}(exps)
		}
	}()

	go func() {
		for {
			pingOrchestrator("agent 007")
			time.Sleep(1 * time.Second)
		}
	}()

	// Wait for program termination
	select {}
}

func processExpression(exp []Expression) {
	// Calculate expression
	var result Result
	for _, expr := range exp {
		result = calculateExpression(expr)
		// Send result to server
		log.Println(result)
		time.Sleep(time.Duration(expr.Cost) * time.Second)
		sendResult(result)
	}
	result.SubExpression_Id = -1
	sendResult(result)
}
