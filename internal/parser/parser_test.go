package parser_test

import (
	"testing"

	"github.com/hellodhlyn/akane/internal/ast/expressions"
	"github.com/hellodhlyn/akane/internal/parser"
	"github.com/stretchr/testify/assert"
)

func assertBinaryExprssion(
	t *testing.T,
	expr expressions.Expression,
	operator string,
	leftType expressions.Expression,
	rightType expressions.Expression,
) *expressions.BinaryExpression {
	assert.IsType(t, &expressions.BinaryExpression{}, expr)

	binExpr := expr.(*expressions.BinaryExpression)
	assert.Equal(t, operator, binExpr.Operator)
	assert.IsType(t, leftType, binExpr.Left)
	assert.IsType(t, rightType, binExpr.Right)

	return binExpr
}

func TestParser_AddExpression(t *testing.T) {
	source := []byte("42 + 9 - -3")
	p := parser.NewParser(source)

	expr, err := p.Parse()
	assert.Nil(t, err)

	binExpr := assertBinaryExprssion(t, expr, "-", &expressions.BinaryExpression{}, &expressions.IntLiteral{})

	leftExpr := assertBinaryExprssion(t, binExpr.Left, "+", &expressions.IntLiteral{}, &expressions.IntLiteral{})
	assert.Equal(t, 42, leftExpr.Left.(*expressions.IntLiteral).Value)
	assert.Equal(t, 9, leftExpr.Right.(*expressions.IntLiteral).Value)

	assert.Equal(t, -3, binExpr.Right.(*expressions.IntLiteral).Value)
}

func TestParser_MulExpression(t *testing.T) {
	source := []byte("42 * 9 / -3")
	p := parser.NewParser(source)

	expr, err := p.Parse()
	assert.Nil(t, err)

	binExpr := assertBinaryExprssion(t, expr, "/", &expressions.BinaryExpression{}, &expressions.IntLiteral{})

	leftExpr := assertBinaryExprssion(t, binExpr.Left, "*", &expressions.IntLiteral{}, &expressions.IntLiteral{})
	assert.Equal(t, 42, leftExpr.Left.(*expressions.IntLiteral).Value)
	assert.Equal(t, 9, leftExpr.Right.(*expressions.IntLiteral).Value)

	assert.Equal(t, -3, binExpr.Right.(*expressions.IntLiteral).Value)
}

func TestParser_ComplexAddMulExpression(t *testing.T) {
	source := []byte("42 + 9 / -3")
	p := parser.NewParser(source)

	expr, err := p.Parse()
	assert.Nil(t, err)

	// 42 + (9 / 3)
	binExpr := assertBinaryExprssion(t, expr, "+", &expressions.IntLiteral{}, &expressions.BinaryExpression{})
	assert.Equal(t, 42, binExpr.Left.(*expressions.IntLiteral).Value)

	assertBinaryExprssion(t, binExpr.Right, "/", &expressions.IntLiteral{}, &expressions.IntLiteral{})
	assert.Equal(t, 9, binExpr.Right.(*expressions.BinaryExpression).Left.(*expressions.IntLiteral).Value)
	assert.Equal(t, -3, binExpr.Right.(*expressions.BinaryExpression).Right.(*expressions.IntLiteral).Value)
}
