package main

import (
	"fmt"
	"golite/cmdline"
	"golite/lexer"
	"golite/parser"
)

func main() {
	// Parse the command line
	cmd := cmdline.ReadCmdline()
	if cmd == nil {
		return
	}

	// Get the input source file path
	// Input source path is always the last argument
	inputSourcePath := cmd.Args[len(cmd.Args)-1]

	// Get a new lexer
	lexer := lexer.NewLexer(inputSourcePath)

	// Get a new parser
	parser := parser.NewParser(lexer)

	// Parse the input source file (lexes and parses)
	invalidParse := parser.Parse()
	if invalidParse {
		fmt.Println("Program as is incorrect!")
	} else {
		fmt.Println("Program is correct! (till now...)")
	}
}
