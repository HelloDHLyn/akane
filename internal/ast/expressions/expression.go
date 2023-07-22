package expressions

import "github.com/hellodhlyn/akane/internal/objects"

type ExpressionType int

const (
	BinaryExpressionType ExpressionType = iota
	IntLiteralType
)

type Expression interface {
	Eval() objects.Object
	Type() ExpressionType
}
