package expressions_test

import (
	"testing"

	"github.com/hellodhlyn/akane/internal/ast/expressions"
	"github.com/hellodhlyn/akane/internal/objects"
	"github.com/stretchr/testify/assert"
)

func TestBinaryExpression_EvalInt(t *testing.T) {
	type Case struct {
		operator []byte
		left     []byte
		right    []byte
		expected int
	}

	cases := []Case{
		{[]byte("+"), []byte("1"), []byte("2"), 3},
		{[]byte("-"), []byte("1"), []byte("2"), -1},
		{[]byte("*"), []byte("2"), []byte("3"), 6},
		{[]byte("/"), []byte("6"), []byte("3"), 2},
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
