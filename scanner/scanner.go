package scanner

import (
	"GoLox/parseError"
	"GoLox/token"
)

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
	case "*":
		scanner.addToken(token.STAR)
	case "!":
		if scanner.match('=') {
			scanner.addToken(token.BANG_EQUAL)
		} else {
			scanner.addToken(token.BANG)
		}
	case "=":
		if scanner.match('=') {
			scanner.addToken(token.EQUAL_EQUAL)
		} else {
			scanner.addToken(token.EQUAL)
		}
	case "<":
		if scanner.match('=') {
			scanner.addToken(token.LESS_EQUAL)
		} else {
			scanner.addToken(token.LESS)
		}
	case ">":
		if scanner.match('=') {
			scanner.addToken(token.GREATER_EQUAL)
		} else {
			scanner.addToken(token.GREATER)
		}
	case "/":
		if scanner.match('/') {
			for scanner.peek() != "\n" && !scanner.isAtEnd() {
				scanner.advance()
			}
		} else {
			scanner.addToken(token.SLASH)
		}
	case " ":

	case "\r":
	case "\t":
		// Ignore whitespace.
		break

	case "\n":
		scanner.line++
	case "\"":
		scanner.addString()
	default:
		parseError.RaiseError(scanner.line, "Unexpected character.")
	}
}
func (scanner *Scanner) advance() string {
	char := string(scanner.source[scanner.current])
	scanner.current++
	return char
}
func (scanner *Scanner) addToken(tp token.Type) {
	scanner.addTokenWithLiteral(tp, nil)
}
func (scanner *Scanner) addTokenWithLiteral(tp token.Type, literal interface{}) {
	text := scanner.source[scanner.start:scanner.current]
	scanner.tokens = append(scanner.tokens, token.Token{Type: tp, Lexeme: text, Literal: literal, Line: scanner.line})
}
func (scanner *Scanner) match(expected byte) bool {
	if scanner.isAtEnd() {
		return false
	}
	if scanner.source[scanner.current] != expected {
		return false
	}
	scanner.current++
	return true
}
func (scanner *Scanner) peek() string {
	if scanner.isAtEnd() {
		return "\\0"
	}
	return string(scanner.source[scanner.current])
}
func (scanner *Scanner) addString() {
	for scanner.peek() != "\"" && !scanner.isAtEnd() {
		if scanner.peek() == "\n" {
			scanner.line++
		}
		scanner.advance()
	}
	if scanner.isAtEnd() {
		parseError.RaiseError(scanner.line, "Unterminated string.")
		return
	}
	scanner.advance()
	value := scanner.source[scanner.start+1 : scanner.current-1]
	scanner.addTokenWithLiteral(token.STRING, value)
}
