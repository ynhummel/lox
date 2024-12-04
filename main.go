package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

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
	}
}

func run(source string) {
	fmt.Print(source)
}
