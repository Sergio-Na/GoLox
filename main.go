package main

import (
	"fmt"
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
