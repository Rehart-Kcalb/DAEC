package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.Handle("POST /api/expression", hand.AddExpression)
	mux.Handle("GET /api/expressions", hand.GetExpressions)
	mux.Handle("GET /api/expressions/{id}", hand.GetExpression)
	mux.Handle("GET /api/operations/", hand.GetOperations)
	mux.Handle("GET /api/task", hand.GiveTask)
	mux.Handle("POST /api/task", hand.GetResult)
	mux.Handle("POST /api/ping", hand.ProcessPing)

	http.ListenAndServe(":8080", mux)
}
