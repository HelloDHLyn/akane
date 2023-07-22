package expressions_test

import (
	"testing"

	"github.com/hellodhlyn/akane/internal/ast/expressions"
	"github.com/hellodhlyn/akane/internal/objects"
	"github.com/stretchr/testify/assert"
)

func TestBinaryExpression_Rotate(t *testing.T) {
	// 42 + 9 - 3
	rightExpr := expressions.NewBinaryExpression("-", expressions.NewIntLiteral([]byte("9")), expressions.NewIntLiteral([]byte("3")))
	expr := expressions.NewBinaryExpression("+", expressions.NewIntLiteral([]byte("42")), rightExpr)
	expr.Rotate()

	assert.Equal(t, "-", expr.Operator)
	assert.IsType(t, &expressions.BinaryExpression{}, expr.Left)
	assert.IsType(t, &expressions.IntLiteral{}, expr.Right)

	leftExpr := expr.Left.(*expressions.BinaryExpression)
	assert.Equal(t, "+", leftExpr.Operator)
	assert.Equal(t, 42, leftExpr.Left.(*expressions.IntLiteral).Value)
	assert.Equal(t, 9, leftExpr.Right.(*expressions.IntLiteral).Value)
}

func TestBinaryExpression_EvalInt(t *testing.T) {
	type Case struct {
		operator string
		left     []byte
		right    []byte
		expected int
	}

	cases := []Case{
		{"+", []byte("1"), []byte("2"), 3},
		{"-", []byte("1"), []byte("2"), -1},
		{"*", []byte("2"), []byte("3"), 6},
		{"/", []byte("6"), []byte("3"), 2},
	}

	for _, c := range cases {
		t.Run(string(c.operator), func(t *testing.T) {
			left := expressions.NewIntLiteral(c.left)
			right := expressions.NewIntLiteral(c.right)
			expr := expressions.NewBinaryExpression(c.operator, left, right)

			result := expr.Eval()
			assert.Equal(t, c.expected, result.(*objects.Integer).Value)
		})
	}
}
