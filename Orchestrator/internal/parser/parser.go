package parser

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
	"strings"
	"unicode"
)

var expressionStack *Stack

func ProcessExpression(exprStr string, operators_cost map[string]int) (*Stack, error) {
	expr, err := parser.ParseExpr(exprStr)
	if err != nil {
		return nil, err
	}
	//log.Println(exprStr)
	if containsConsecutiveOperators(exprStr) {
		return nil, errors.New("Error: Invalid expression. Consecutive operators are not allowed.")
	}

	// Create a stack for expressions
	expressionStack = &Stack{}

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
		newExpression.Operand1 = extractOperand(exprStr, binExpr.X)

		// Extract the operator
		newExpression.Operator = binExpr.Op.String()
		newExpression.Cost = operators_cost[newExpression.Operator]

		// Extract the right operand
		newExpression.Operand2 = extractOperand(exprStr, binExpr.Y)

		// Push the new expression onto the stack
		expressionStack.Push(*newExpression)

		// Stop traversal
		return true
	})

	/*for i, val := range expressionStack.items {*/
	/*expressionStack.items[i].Operand1 = FindExp(strings.ReplaceAll(val.Operand1, " ", ""))*/
	/*expressionStack.items[i].Operand2 = FindExp(strings.ReplaceAll(val.Operand2, " ", ""))*/
	/*}*/
	/*log.Println(expressionStack.items)*/

	return expressionStack, nil
}

// extractOperand extracts the operand from the expression string based on the position information in the AST node.
func extractOperand(exprStr string, expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.BasicLit:
		// If the expression is a basic literal (e.g., a number), return its value
		if t.Kind == token.INT {
			return t.Value + ".0"
		}
		return t.Value
	case *ast.ParenExpr:
		// If the expression is a parenthesized expression, recursively extract the operand inside the parentheses
		return extractOperand(exprStr, t.X)
	default:
		// If the expression is not a basic literal or a parenthesized expression, return the substring from the expression string
		pos := expr.Pos() - 1
		end := expr.End() - 1
		//return "$" + FindExp(strings.ReplaceAll(exprStr[pos:end], " ", ""))
		return convertToFloat(exprStr[pos:end])
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

/*func FindExp(exprStr string) string {*/
/*//log.Println(expressionStack.items)*/
/*reverse(expressionStack.items)*/
/*for i, exp := range expressionStack.items {*/
/*log.Printf("%v\n%v", exprStr, (exp.Operand1 + exp.Operator + exp.Operand2))*/
/*if exprStr == (exp.Operand1 + exp.Operator + exp.Operand2) {*/
/*return fmt.Sprintf("$%d", i)*/
/*}*/
/*}*/
/*return exprStr*/
/*}*/

/*func reverse(s []Expression) {*/
/*for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {*/
/*s[i], s[j] = s[j], s[i]*/
/*}*/
/*}*/

func convertToFloat(expression string) string {
	// Split the expression into individual tokens
	tokens := make([]string, 0)
	token := ""
	for _, char := range expression {
		if unicode.IsDigit(char) || char == '.' {
			// If the character is a digit or '.', add it to the current token
			token += string(char)
		} else {
			// If the character is an operator, add the current token (if not empty)
			// to the list of tokens and reset the token variable
			if token != "" {
				tokens = append(tokens, token)
				token = ""
			}
			// Add the operator to the list of tokens
			tokens = append(tokens, string(char))
		}
	}
	// Add the last token (if not empty) to the list of tokens
	if token != "" {
		tokens = append(tokens, token)
	}

	// Process each token
	for i, token := range tokens {
		// Attempt to convert the token to a floating point number
		if num, err := strconv.ParseFloat(token, 64); err == nil {
			// Replace the token with its floating point representation
			tokens[i] = fmt.Sprintf("%.1f", num)
		}
	}

	// Reconstruct the expression with converted operands
	result := strings.Join(tokens, "")
	return result
}
