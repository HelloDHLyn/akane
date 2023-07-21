package lexer_test

import (
	"testing"

	"github.com/hellodhlyn/akane/internal/lexer"
	"github.com/stretchr/testify/assert"
)

func TestScanner_Int(t *testing.T) {
	type Case struct {
		source []byte
		kind   lexer.TokenKind
		lexeme []byte
	}

	cases := []Case{
		{[]byte("123"), lexer.TokenIntLiteral, []byte("123")},
		{[]byte("-123"), lexer.TokenIntLiteral, []byte("-123")},
	}

	for _, c := range cases {
		scanner := lexer.NewScanner(c.source)
		token := scanner.Scan()

		assert.Equal(t, c.kind, token.Kind)
		assert.Equal(t, c.lexeme, token.Lexeme)
	}
}

func TestScanner_ArithmeticOperators(t *testing.T) {
	type Case struct {
		source []byte
		kind   lexer.TokenKind
	}

	cases := []Case{
		{[]byte("+"), lexer.TokenAdd},
		{[]byte("-"), lexer.TokenSub},
		{[]byte("*"), lexer.TokenMul},
		{[]byte("/"), lexer.TokenDiv},
	}

	for _, c := range cases {
		scanner := lexer.NewScanner(c.source)
		token := scanner.Scan()

		assert.Equal(t, c.kind, token.Kind)
		assert.Equal(t, c.source, token.Lexeme)
	}
}

func TestScanner_ArithmeticOperations(t *testing.T) {
	type Case struct {
		source []byte
		kinds  []lexer.TokenKind
	}

	cases := []Case{
		{[]byte("42+9"), []lexer.TokenKind{lexer.TokenIntLiteral, lexer.TokenAdd, lexer.TokenIntLiteral}},
		{[]byte("42 * -9"), []lexer.TokenKind{lexer.TokenIntLiteral, lexer.TokenMul, lexer.TokenIntLiteral}},
	}

	for _, c := range cases {
		scanner := lexer.NewScanner(c.source)
		for _, kind := range c.kinds {
			token := scanner.Scan()
			assert.Equal(t, kind, token.Kind)
		}
	}
}
