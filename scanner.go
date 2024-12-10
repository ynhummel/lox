package main

import "fmt"

type Scanner struct {
	Source string
	Tokens []Token

	start, current, line int
}

func newScanner(source string) *Scanner {
	return &Scanner{Source: source, start: 0, current: 0, line: 1}
}

func (scn *Scanner) ScanTokens() {
	for !scn.atEnd() {
		scn.start = scn.current
		scn.scanToken()
	}
}

func (scn *Scanner) atEnd() bool {
	return scn.current >= len(scn.Source)
}

func (scn *Scanner) scanToken() {
	rn := scn.advance()
	switch rn {
	case '(':
		scn.addToken(LEFT_PAREN)
	case ')':
		scn.addToken(RIGHT_PAREN)
	case '{':
		scn.addToken(LEFT_BRACE)
	case '}':
		scn.addToken(RIGHT_BRACE)
	case ',':
		scn.addToken(COMMA)
	case '.':
		scn.addToken(DOT)
	case '-':
		scn.addToken(MINUS)
	case '+':
		scn.addToken(PLUS)
	case ';':
		scn.addToken(SEMICOLON)
	case '*':
		scn.addToken(STAR)
	case '!':
		if scn.match('=') {
			scn.addToken(BANG_EQUAL)
		} else {
			scn.addToken(BANG)
		}
	case '=':
		if scn.match('=') {
			scn.addToken(EQUAL_EQUAL)
		} else {
			scn.addToken(EQUAL)
		}
	case '<':
		if scn.match('=') {
			scn.addToken(LESS_EQUAL)
		} else {
			scn.addToken(LESS)
		}
	case '>':
		if scn.match('=') {
			scn.addToken(GREATER_EQUAL)
		} else {
			scn.addToken(GREATER)
		}
	default:
		fmt.Println(Err{Line: scn.line, Message: "Unexpected character."}.Error())
	}
}

func (scn *Scanner) match(expected rune) bool {
	if scn.atEnd() || []rune(scn.Source)[scn.current] != expected {
		return false
	}

	scn.current++
	return true
}

func (scn *Scanner) advance() rune {
	oldCurrent := scn.current
	scn.current++
	return []rune(scn.Source)[oldCurrent]
}

func (scn *Scanner) addToken(tokenType TokenType) {
	scn.addTokenWLiteral(tokenType, nil)
}

func (scn *Scanner) addTokenWLiteral(tokenType TokenType, literal any) {
	text := scn.Source[scn.start:scn.current]
	scn.Tokens = append(scn.Tokens, *NewToken(tokenType, text, literal, scn.line))
}
