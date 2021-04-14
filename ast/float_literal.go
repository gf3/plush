package ast

type FloatLiteral struct {
	TokenAble
	Value float64
}

func (il *FloatLiteral) validIfCondition() bool { return true }

func (il *FloatLiteral) expressionNode() {}

func (il *FloatLiteral) String() string {
	return il.Token.Literal
}
