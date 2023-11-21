package main

import "fmt"

type Type string
type Token struct {
	Type    Type
	Lexeme  string
	Literal interface{}
	Line    int
}

const (
	// Single-character tokens
	LEFT_PAREN  = "("
	RIGHT_PAREN = ")"
	LEFT_BRACE  = "["
	RIGHT_BRACE = "]"
	COMMA       = ","
	DOT         = "."
	MINUS       = "-"
	PLUS        = "+"
	SEMICOLON   = ";"
	SLASH       = "/"
	STAR        = "*"
	// One or two character tokens
	BANG          = "!"
	BANG_EQUAL    = "!="
	EQUAL         = "="
	EQUAL_EQUAL   = "=="
	GREATER       = ">"
	GREATER_EQUAL = ">="
	LESS          = "<"
	LESS_EQUAL    = "<="
	// Literals
	IDENTIFIER = "IDENT"
	STRING     = "STRING"
	NUMBER     = "NUMBER"
	// keywords
	AND      = "and"
	CLASS    = "class"
	ELSE     = "else"
	FALSE    = "false"
	FUN      = "fun"
	FOR      = "for"
	IF       = "if"
	NIL      = "nil"
	OR       = "or"
	PRINT    = "print"
	RETURN   = "return"
	SUPER    = "super"
	THIS     = "this"
	TRUE     = "true"
	VAR      = "var"
	WHILE    = "while"
	BREAK    = "break"
	CONTINUE = "continue"
	EOF      = "eof"
	INVALID  = "__INVALID__"
)

func (token *Token) String() string {
	return fmt.Sprintf("%s %s %v", token.Type, token.Lexeme, token.Literal)
}
