package lexer

type Scanner struct {
	source []byte

	currPos    int
	currLexeme []byte
}

func NewScanner(source []byte) *Scanner {
	return &Scanner{
		source: source,
	}
}

func (s *Scanner) Scan() *Token {
	s.currLexeme = []byte{}
	return &Token{
		Kind:   s.scanToken(),
		Lexeme: s.currLexeme,
	}
}

// Character-level operations
func (s *Scanner) currChar() byte {
	if s.currPos >= len(s.source) {
		return 0
	}
	return s.source[s.currPos]
}

func (s *Scanner) dropChar() {
	s.currPos++
}

func (s *Scanner) takeChar() {
	s.currLexeme = append(s.currLexeme, s.currChar())
	s.currPos++
}

// Token-level operations
func (s *Scanner) scanToken() TokenKind {
	for isWhitespace(s.currChar()) || s.currChar() == '\n' {
		s.dropChar()
	}
	if s.currChar() == 0 {
		return TokenEOF
	}

	if s.currChar() == '-' {
		s.takeChar()
		if isDigit(s.currChar()) {
			return s.scanInt()
		} else if isWhitespace(s.currChar()) || isEOL(s.currChar()) {
			return TokenSub
		}
		return TokenErr
	}

	if isDigit(s.currChar()) {
		return s.scanInt()
	}

	switch s.currChar() {
	case '+':
		s.takeChar()
		return TokenAdd
	case '*':
		s.takeChar()
		return TokenMul
	case '/':
		s.takeChar()
		return TokenDiv
	}

	return TokenErr
}

func (s *Scanner) scanInt() TokenKind {
	for isDigit(s.currChar()) {
		s.takeChar()
	}
	return TokenIntLiteral
}

func isWhitespace(c byte) bool {
	return c == ' ' || c == '\t'
}

func isEOL(c byte) bool {
	return c == '\n' || c == 0
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
