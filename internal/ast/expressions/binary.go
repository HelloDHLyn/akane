package expressions

import "github.com/hellodhlyn/akane/internal/objects"

type BinaryExpression struct {
	Operator []byte
	Left     Expression
	Right    Expression
}

func NewBinaryExpression(operator []byte, left, right Expression) *BinaryExpression {
	return &BinaryExpression{
		Operator: operator,
		Left:     left,
		Right:    right,
	}
}

func (expr *BinaryExpression) Eval() objects.Object {
	left := expr.Left.Eval()
	switch left.Type() {
	case objects.IntegerObject:
		return expr.evalInt(left.(*objects.Integer))
	default: // TODO
	}
	return nil
}

func (expr *BinaryExpression) evalInt(leftInt *objects.Integer) objects.Object {
	right := expr.Right.Eval()
	switch right.Type() {
	case objects.IntegerObject:
		rightInt := right.(*objects.Integer)
		switch expr.Operator[0] {
		case '+':
			return leftInt.Add(rightInt)
		case '-':
			return leftInt.Sub(rightInt)
		case '*':
			return leftInt.Mul(rightInt)
		case '/':
			return leftInt.Div(rightInt)
		default: // TODO
		}
	default: // TODO
	}
	return nil
}
