package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/haritsrizkall/golox/scanner"
)

var hadError = false

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Println("Usage:golox [script]")
	} else if len(args) == 1 {
		runFile(args[0])
	} else {
		runPrompt()
	}
}

func error(line int, message string) {
	report(line, "", message)
}

func report(line int, where, message string) {
	fmt.Errorf("[line %d] Error %s: %s", line, where, message)
	hadError = true
}

func run(str string) {
	fmt.Printf("Run %s\n", str)
	scanner := &scanner.Scanner{}
	tokens := scanner.ScanTokens()
	fmt.Println("Scan tokens")
	for _, token := range tokens {
		fmt.Println(token)
	}

}

func runPrompt() {
	fmt.Println("Runprompt")
	reader := bufio.NewReader(os.Stdin)
	for {
		var line string
		fmt.Print("> ")
		reader.ReadString('\n')
		run(line)
		if hadError {
			os.Exit(65)
		}
		hadError = false
	}
}

func runFile(path string) {
	fmt.Println("Runfile")
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	run(string(bytes))
	if hadError {
		os.Exit(65)
	}
}
