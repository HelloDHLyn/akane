package parser

import (
	"errors"

	"github.com/hellodhlyn/akane/internal/ast/expressions"
	"github.com/hellodhlyn/akane/internal/lexer"
)

type Parser struct {
	lexer     *lexer.Scanner
	currToken *lexer.Token
}

func NewParser(source []byte) *Parser {
	lexer := lexer.NewScanner(source)
	return &Parser{
		lexer:     lexer,
		currToken: lexer.Scan(),
	}
}

func (p *Parser) Parse() (expressions.Expression, error) {
	expr := p.parseAddExpression()
	if expr == nil {
		return nil, errors.New("failed to parse experssion")
	}
	return expr, nil
}

func (p *Parser) takeToken() *lexer.Token {
	token := p.currToken
	p.currToken = p.lexer.Scan()
	return token
}

// add_expr -> mul_expr (('+'|'-') add_expr)*
func (p *Parser) parseAddExpression() expressions.Expression {
	expr := p.parseMulExpression()
	if expr == nil {
		return nil
	}

	for p.currToken.Kind == lexer.TokenAdd || p.currToken.Kind == lexer.TokenSub {
		operator := p.takeToken()
		right := p.parseAddExpression()
		if right == nil {
			return nil
		}

		expr = expressions.NewBinaryExpression(operator.LexemeString(), expr, right)
		if expr.(*expressions.BinaryExpression).Right.Type() == expressions.BinaryExpressionType {
			expr.(*expressions.BinaryExpression).Rotate()
		}
	}
	return expr
}

// mul_expr -> IntLiteral (('*'|'/') mul_expr)*
func (p *Parser) parseMulExpression() expressions.Expression {
	expr := p.parseIntLiteral()
	if expr == nil {
		return nil
	}

	for p.currToken.Kind == lexer.TokenMul || p.currToken.Kind == lexer.TokenDiv {
		operator := p.takeToken()
		right := p.parseMulExpression()
		if right == nil {
			return nil
		}

		expr = expressions.NewBinaryExpression(operator.LexemeString(), expr, right)
		if expr.(*expressions.BinaryExpression).Right.Type() == expressions.BinaryExpressionType {
			expr.(*expressions.BinaryExpression).Rotate()
		}
	}
	return expr
}

// IntLiteral
func (p *Parser) parseIntLiteral() expressions.Expression {
	if p.currToken.Kind != lexer.TokenIntLiteral {
		return nil
	}
	token := p.takeToken()
	return expressions.NewIntLiteral(token.Lexeme)
}
