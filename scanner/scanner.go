package scanner

import "GoLox/token"

type Scanner struct {
	source  string
	tokens  []token.Token
	start   int
	current int
	line    int
}

func New(source string) Scanner {
	scanner := Scanner{source: source, tokens: make([]token.Token, 0), start: 0, current: 0, line: 1}
	return scanner
}
func (scanner *Scanner) setSource(source string) {
	scanner.source = source
}
func (scanner *Scanner) scanTokens() {
	for scanner.isAtEnd() {
		scanner.start = scanner.current
		scanner.scanToken()
	}
	scanner.tokens = append(scanner.tokens, token.Token{Type: token.EOF, Lexeme: "", Literal: nil, Line: scanner.line})
}
func (scanner *Scanner) isAtEnd() bool {
	return scanner.current >= len(scanner.source)
}
func (scanner *Scanner) scanToken() {
	char := scanner.advance()
	switch char {
	case '(':
	}

}
func (scanner *Scanner) advance() string {
	char := string(scanner.source[scanner.current])
	scanner.current++
	return char
}
func (scanner *Scanner) addToken(t token.Token) {
	scanner.tokens = append(scanner.tokens, t)
}
