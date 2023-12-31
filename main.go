package main

import (
	"GoLox/parseError"
	"bufio"
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

func runFile(filename string) {
	file_bytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read file: %v", err)
	}

	run(string(file_bytes))

	if parseError.HadError {
		os.Exit(65)
	}
}
func runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Print("> ")
		line := scanner.Text()
		run(line)
		parseError.HadError = false
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read next line in REPL")
	}
}

func run(source string) {
	tokens, err := getTokens(source)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run source")
	}
	for i := 0; i < len(tokens); i++ {
		fmt.Println(tokens[i])
	}

}

func getTokens(source string) ([]string, error) {
	return []string{}, nil
}
