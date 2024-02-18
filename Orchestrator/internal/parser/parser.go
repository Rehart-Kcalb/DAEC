package parser

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"strings"
)

// Expression represents a mathematical expression.

func processExpression(exprStr string) error {
	expr, err := parser.ParseExpr(exprStr)
	if err != nil {
		return err
	}
	fmt.Print(expr)
	// Check for consecutive operators
	if containsConsecutiveOperators(exprStr) {
		return errors.New("Error: Invalid expression. Consecutive operators are not allowed.")
	}

	// Create a stack for expressions
	expressionStack := &Stack{}

	// Traverse the expression tree
	ast.Inspect(expr, func(n ast.Node) bool {
		if n == nil {
			return false
		}

		// Check if the node is a binary expression
		binExpr, ok := n.(*ast.BinaryExpr)
		if !ok {
			return true
		}

		// Create a new Expression struct
		newExpression := &Expression{}

		// Extract the left operand
		newExpression.LeftOperand = extractOperand(exprStr, binExpr.X)

		// Extract the operator
		newExpression.Operator = binExpr.Op.String()

		// Extract the right operand
		newExpression.RightOperand = extractOperand(exprStr, binExpr.Y)

		// Push the new expression onto the stack
		expressionStack.Push(*newExpression)

		// Stop traversal
		return true
	})

	// Evaluate the expressions in the stack
	for !expressionStack.IsEmpty() {
		fmt.Println(expressionStack.Pop())
	}
	return nil
}

// extractOperand extracts the operand from the expression string based on the position information in the AST node.
func extractOperand(exprStr string, expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.BasicLit:
		// If the expression is a basic literal (e.g., a number), return its value
		return t.Value
	case *ast.ParenExpr:
		// If the expression is a parenthesized expression, recursively extract the operand inside the parentheses
		return extractOperand(exprStr, t.X)
	default:
		// If the expression is not a basic literal or a parenthesized expression, return the substring from the expression string
		pos := expr.Pos() - 1
		end := expr.End() - 1
		return exprStr[pos:end]
	}
}

func containsConsecutiveOperators(exprStr string) bool {
	ops := "+-*/"
	for i := 0; i < len(exprStr)-1; i++ {
		if strings.Contains(ops, string(exprStr[i])) && strings.Contains(ops, string(exprStr[i+1])) {
			return true
		}
	}
	return false
}
