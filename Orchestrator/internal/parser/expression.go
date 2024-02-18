package parser

type Expression struct {
	LeftOperand   string
	Operator      string
	RightOperand  string
	Result        int
	DependencyPtr *Expression
}
