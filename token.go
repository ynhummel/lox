package main

import "fmt"

type Token struct {
	Type    TokenType
	lexeme  string
	literal any
	line    int
}

func NewToken(tokenType TokenType, lexeme string, literal any, line int) *Token {
	return &Token{tokenType, lexeme, literal, line}
}

func (t Token) String() string {
	return fmt.Sprintf("%v %s %v", t.Type, t.lexeme, t.literal)
}
