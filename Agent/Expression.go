package main

type Expression struct {
	Id            int64  `json:"id"`
	Expression_Id int64  `json:"expression_id"`
	Operation     string `json:"operation"`
	Operand1      string `json:"operand1"`
	Operand2      string `json:"operand2"`
	ExecOrder     int    `json:"exec_order"`
	Cost          int    `json:"cost"`
}

type Result struct {
	Expression_Id    int64   `json:"expression_id"`
	SubExpression_Id int64   `json:"sub_expression_id"`
	Value            float64 `json:"result"`
}
