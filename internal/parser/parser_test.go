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

	world, err := p.Parse()
	assert.Nil(t, err)

	expr := world.Expressions[0]
	binExpr := assertBinaryExprssion(t, expr, "-", &expressions.BinaryExpression{}, &expressions.IntLiteral{})

	leftExpr := assertBinaryExprssion(t, binExpr.Left, "+", &expressions.IntLiteral{}, &expressions.IntLiteral{})
	assert.Equal(t, 42, leftExpr.Left.(*expressions.IntLiteral).Value)
	assert.Equal(t, 9, leftExpr.Right.(*expressions.IntLiteral).Value)

	assert.Equal(t, -3, binExpr.Right.(*expressions.IntLiteral).Value)
}

func TestParser_MulExpression(t *testing.T) {
	source := []byte("42 * 9 / -3")
	p := parser.NewParser(source)

	world, err := p.Parse()
	assert.Nil(t, err)

	expr := world.Expressions[0]
	binExpr := assertBinaryExprssion(t, expr, "/", &expressions.BinaryExpression{}, &expressions.IntLiteral{})

	leftExpr := assertBinaryExprssion(t, binExpr.Left, "*", &expressions.IntLiteral{}, &expressions.IntLiteral{})
	assert.Equal(t, 42, leftExpr.Left.(*expressions.IntLiteral).Value)
	assert.Equal(t, 9, leftExpr.Right.(*expressions.IntLiteral).Value)

	assert.Equal(t, -3, binExpr.Right.(*expressions.IntLiteral).Value)
}

func TestParser_ComplexAddMulExpression(t *testing.T) {
	source := []byte("42 + 9 / -3")
	p := parser.NewParser(source)

	world, err := p.Parse()
	assert.Nil(t, err)

	// 42 + (9 / 3)
	expr := world.Expressions[0]
	binExpr := assertBinaryExprssion(t, expr, "+", &expressions.IntLiteral{}, &expressions.BinaryExpression{})
	assert.Equal(t, 42, binExpr.Left.(*expressions.IntLiteral).Value)

	assertBinaryExprssion(t, binExpr.Right, "/", &expressions.IntLiteral{}, &expressions.IntLiteral{})
	assert.Equal(t, 9, binExpr.Right.(*expressions.BinaryExpression).Left.(*expressions.IntLiteral).Value)
	assert.Equal(t, -3, binExpr.Right.(*expressions.BinaryExpression).Right.(*expressions.IntLiteral).Value)
}

func TestParser_MultipleExpressions(t *testing.T) {
	source := []byte("42 + 9\n9 * 42")
	p := parser.NewParser(source)

	world, err := p.Parse()
	assert.Nil(t, err)

	assert.Equal(t, 2, len(world.Expressions))

	firstExpr := world.Expressions[0]
	assertBinaryExprssion(t, firstExpr, "+", &expressions.IntLiteral{}, &expressions.IntLiteral{})
	assert.Equal(t, 42, firstExpr.(*expressions.BinaryExpression).Left.(*expressions.IntLiteral).Value)
	assert.Equal(t, 9, firstExpr.(*expressions.BinaryExpression).Right.(*expressions.IntLiteral).Value)

	secondExpr := world.Expressions[1]
	assertBinaryExprssion(t, secondExpr, "*", &expressions.IntLiteral{}, &expressions.IntLiteral{})
	assert.Equal(t, 9, secondExpr.(*expressions.BinaryExpression).Left.(*expressions.IntLiteral).Value)
	assert.Equal(t, 42, secondExpr.(*expressions.BinaryExpression).Right.(*expressions.IntLiteral).Value)
}
