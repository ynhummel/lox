package scanner

import (
	"fmt"
	"strconv"

	interror "github.com/ynhummel/lox/error"
)

type Scanner struct {
	Source string
	Tokens []Token

	start   int
	current int
	line    int
}

func NewScanner(source string) *Scanner {
	return &Scanner{Source: source, start: 0, current: 0, line: 1}
}

func (scn *Scanner) ScanTokens() []Token {
	for !scn.atEnd() {
		scn.start = scn.current
		scn.scanToken()
	}

	scn.Tokens = append(scn.Tokens, *NewToken(EOF, "", nil, scn.line))
	return scn.Tokens
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
	case '/':
		if scn.match('/') {
			for scn.peek() != '\n' && !scn.atEnd() {
				scn.advance()
			}
		} else {
			scn.addToken(SLASH)
		}
	case '"':
		scn.scanString()
	case ' ', '\r', '\t':
	case '\n':
		scn.line++
	default:
		if isDigit(rn) {
			scn.scanNumber()
		} else if isAlpha(rn) {
			scn.scanIdentifier()
		} else {
			fmt.Println(interror.Err{Line: scn.line, Message: "Unexpected character."}.Error())
		}
	}
}

func (scn *Scanner) addToken(tokenType TokenType) {
	scn.addTokenWLiteral(tokenType, nil)
}

func (scn *Scanner) addTokenWLiteral(tokenType TokenType, literal any) {
	text := scn.Source[scn.start:scn.current]
	scn.Tokens = append(scn.Tokens, *NewToken(tokenType, text, literal, scn.line))
}

func (scn *Scanner) advance() rune {
	oldCurrent := scn.current
	scn.current++
	return []rune(scn.Source)[oldCurrent]
}

func (scn *Scanner) match(expected rune) bool {
	if scn.atEnd() || []rune(scn.Source)[scn.current] != expected {
		return false
	}

	scn.current++
	return true
}
func (scn *Scanner) scanIdentifier() {
	for isAlphanumeric(scn.peek()) {
		scn.advance()
	}

	text := []rune(scn.Source)[scn.start:scn.current]
	if tokenType, ok := Keywords[string(text)]; ok {
		scn.addToken(tokenType)
	} else {
		scn.addToken(IDENTIFIER)
	}
}

func isAlphanumeric(rn rune) bool {
	return isAlpha(rn) || isDigit(rn)
}

func isAlpha(rn rune) bool {
	return (rn >= 'a' && rn <= 'z') || (rn >= 'A' && rn <= 'Z') || rn == '_'
}

func (scn *Scanner) scanString() {
	for scn.peek() != '"' && !scn.atEnd() {
		if scn.peek() == '\n' {
			scn.line++
		}
		scn.advance()
	}

	if scn.atEnd() {
		fmt.Println(interror.Err{Line: scn.line, Message: "Unexpected character."}.Error())
		return
	}

	scn.advance()
	scn.addTokenWLiteral(STRING, []rune(scn.Source)[scn.start:scn.current-1])
}

func (scn *Scanner) scanNumber() {
	for isDigit(scn.peek()) {
		scn.advance()
	}

	if scn.peek() == '.' && isDigit(scn.peekNext()) {
		scn.advance()

		for isDigit(scn.peek()) {
			scn.advance()
		}
	}

	value, err := strconv.ParseFloat(string([]rune(scn.Source)[scn.start:scn.current]), 64)
	if err != nil {
		panic(err)
	}
	scn.addTokenWLiteral(NUMBER, value)
}

func isDigit(rn rune) bool {
	return rn >= '0' && rn <= '9'
}

func (scn *Scanner) peek() rune {
	if scn.atEnd() {
		return '\x00'
	}
	return []rune(scn.Source)[scn.current]

}

func (scn *Scanner) peekNext() rune {
	if scn.current+1 >= len(scn.Source) {
		return '\x00'
	}
	return []rune(scn.Source)[scn.current+1]
}
