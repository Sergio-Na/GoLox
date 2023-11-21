package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func runFile(filename string) {
	file_bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	run(string(file_bytes))
}
func runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Print("> ")
		line := scanner.Text()
		run(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Failed to read next line in REPL")
	}
}

func run(source string) {

}

func getTokens() {

}
