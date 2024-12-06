package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
}

type Err struct {
	Message string
	Where   string
	Line    int
}

func (e Err) Error() string {
	HadError = true
	return fmt.Sprintf("[line %d] Error%s: %s", e.Line, e.Where, e.Message)
}
