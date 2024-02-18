package main

import (
	"github.com/expr-lang/expr"
)

func eval(exp string) (float64, error) {
	program, err := expr.Compile(exp)
	if err != nil {
		return 0, err
	}
	output, err := expr.Run(program, nil)
	if err != nil {
		return 0, err
	}

	return float64(output.(float64)), nil
}
