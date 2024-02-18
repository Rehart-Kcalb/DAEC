package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func getExpressionFromOrchestrator(url string, calculator_id int) ([]Expression, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?calculator=%d", url, calculator_id), nil)
	if err != nil {
		log.Println("Error, while creating request", err)
	}
	req.Header.Add("X-Agent-Name", "Agent"+req.RemoteAddr)
	resp, err := client.Do(req)
	if err != nil {
		return []Expression{}, err
	}
	defer resp.Body.Close()

	var exp struct {
		Status int `json:"status"`
		Data   struct {
			Expression    struct{}     `json:"expression"`
			SubExpression []Expression `json:"sub_expressions"`
		} `json:"data"`
	}
	err = json.NewDecoder(resp.Body).Decode(&exp)
	if err != nil {
		return []Expression{}, err
	}
	buf, _ := json.Marshal(exp.Data)
	log.Println(string(buf))
	return exp.Data.SubExpression, nil
	//return exp.Data, nil
}
