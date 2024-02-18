package parser

type Stack struct {
	items []Expression
}

// Push adds an Expression to the top of the stack
func (s *Stack) Push(expr Expression) {
	s.items = append(s.items, expr)
}

// Pop removes and returns the top Expression from the stack
func (s *Stack) Pop() (Expression, bool) {
	if len(s.items) == 0 {
		return Expression{}, false
	}
	lastIndex := len(s.items) - 1
	expr := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return expr, true
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Peek returns the top Expression without removing it
func (s *Stack) Peek() (Expression, bool) {
	if len(s.items) == 0 {
		return Expression{}, false
	}
	lastIndex := len(s.items) - 1
	return s.items[lastIndex], true
}
