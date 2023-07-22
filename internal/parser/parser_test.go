package parser_test

import (
	"testing"

	"github.com/hellodhlyn/akane/internal/ast/expressions"
	"github.com/hellodhlyn/akane/internal/parser"
	"github.com/stretchr/testify/assert"
)

func TestParser_AddExpression(t *testing.T) {
	source := []byte("42 + 9 - -3")
	p := parser.NewParser(source)
	expr, err := p.Parse()

	assert.Nil(t, err)
	assert.IsType(t, &expressions.BinaryExpression{}, expr)

	// (42 + 9) - 3
	binExpr := expr.(*expressions.BinaryExpression)
	assert.Equal(t, []byte("-"), binExpr.Operator)
	assert.IsType(t, &expressions.BinaryExpression{}, binExpr.Left)
	assert.IsType(t, &expressions.IntLiteral{}, binExpr.Right)

	// 42 + 9
	leftExpr := binExpr.Left.(*expressions.BinaryExpression)
	assert.Equal(t, []byte("+"), leftExpr.Operator)
	assert.IsType(t, &expressions.IntLiteral{}, leftExpr.Left)
	assert.IsType(t, &expressions.IntLiteral{}, leftExpr.Right)
	assert.Equal(t, 42, leftExpr.Left.(*expressions.IntLiteral).Value)
	assert.Equal(t, 9, leftExpr.Right.(*expressions.IntLiteral).Value)

	// 3
	assert.Equal(t, -3, binExpr.Right.(*expressions.IntLiteral).Value)
}
