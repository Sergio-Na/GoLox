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
	case "(":
		scanner.addToken(token.LEFT_PAREN)
	case ")":
		scanner.addToken(token.LEFT_PAREN)
	case "{}":
		scanner.addToken(token.LEFT_BRACE)
	case "}":
		scanner.addToken(token.RIGHT_BRACE)
	case ".":
		scanner.addToken(token.DOT)
	case ",":
		scanner.addToken(token.COMMA)
	case "-":
		scanner.addToken(token.MINUS)
	case "+":
		scanner.addToken(token.PLUS)
	case ";":
		scanner.addToken(token.SEMICOLON)
	case "/":
		scanner.addToken(token.SLASH)
	case "*":
		scanner.addToken(token.STAR)

	}

}
func (scanner *Scanner) advance() string {
	char := string(scanner.source[scanner.current])
	scanner.current++
	return char
}
func (scanner *Scanner) addToken(t token.Type) {
	scanner.tokens = append(scanner.tokens, token.Token{Type: t})
}
