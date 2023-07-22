package lexer

// TokenKind is an enum of all the kinds of tokens that the lexer can produce
type TokenKind int

const (
	TokenIntLiteral TokenKind = iota // 123

	// Arithmetic operators
	TokenAdd // +
	TokenSub // -
	TokenMul // *
	TokenDiv // /

	TokenErr // Unexpected character
)

type Token struct {
	Kind   TokenKind
	Lexeme []byte
}

func (t *Token) LexemeString() string {
	return string(t.Lexeme)
}
