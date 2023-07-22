package expressions

import "github.com/hellodhlyn/akane/internal/objects"

var precendence = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}

type BinaryExpression struct {
	Operator string
	Left     Expression
	Right    Expression
}

func NewBinaryExpression(operator string, left, right Expression) *BinaryExpression {
	return &BinaryExpression{
		Operator: operator,
		Left:     left,
		Right:    right,
	}
}

func (*BinaryExpression) Type() ExpressionType {
	return BinaryExpressionType
}

func (expr *BinaryExpression) Rotate() {
	rightExpr := expr.Right.(*BinaryExpression)
	if precendence[expr.Operator] < precendence[rightExpr.Operator] {
		return
	}

	rightExpr.Operator, expr.Operator = expr.Operator, rightExpr.Operator
	rightExpr.Left, rightExpr.Right, expr.Left, expr.Right = expr.Left, rightExpr.Left, expr.Right, rightExpr.Right
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
