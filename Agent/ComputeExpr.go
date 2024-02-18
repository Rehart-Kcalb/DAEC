package main

import "strconv"

func calculateExpression(exp Expression) Result {
	var result float64

	operand1Value, err := strconv.ParseFloat(exp.Operand1, 64)
	if err != nil {
		operand1Value, err = eval(exp.Operand1)
	}

	operand2Value, err := strconv.ParseFloat(exp.Operand2, 64)
	if err != nil {
		operand2Value, err = eval(exp.Operand2)
	}

	switch exp.Operation {
	case "+":
		result = operand1Value + operand2Value
	case "-":
		result = operand1Value - operand2Value
	case "*":
		result = operand1Value * operand2Value
	case "/":
		if operand2Value == 0 {
			result = 0
		} else {
			result = operand1Value / operand2Value
		}
	}
	return Result{Expression_Id: exp.Expression_Id, Value: result, SubExpression_Id: exp.Id}
}
