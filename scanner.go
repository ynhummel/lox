package main

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

func (scn *Scanner) scanToken() {}
