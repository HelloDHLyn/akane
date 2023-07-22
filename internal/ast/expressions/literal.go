package expressions

import (
	"strconv"

	"github.com/hellodhlyn/akane/internal/objects"
)

type IntLiteral struct {
	Value int
}

func NewIntLiteral(value []byte) *IntLiteral {
	intValue, _ := strconv.Atoi(string(value)) // TODO error handling
	return &IntLiteral{Value: intValue}
}

func (*IntLiteral) Type() ExpressionType {
	return IntLiteralType
}

func (i *IntLiteral) Eval() objects.Object {
	return objects.NewInteger(i.Value)
}
