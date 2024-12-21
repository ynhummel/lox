package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/ynhummel/lox/scanner"
)

var HadError = false

func main() {
	args := os.Args[1:]

	if len(args) > 1 {
		fmt.Println("Usage: jlox [script]")
		os.Exit(64)
	} else if len(args) == 1 {
		runFile(args[0])
	} else {
		runPrompt()
	}
}

func runFile(path string) {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	run(string(file))

	if HadError {
		os.Exit(65)
	}
}

func runPrompt() {
	scn := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		line, err := scn.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		run(line)
		HadError = false

	}
}

func run(source string) {
	fmt.Print(source)
	scn := scanner.NewScanner(source)
	tokens := scn.ScanTokens()

	for _, tkn := range tokens {
		fmt.Printf("%+v\n", tkn)
	}
}
