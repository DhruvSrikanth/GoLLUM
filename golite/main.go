package main

import (
	"fmt"
	"golite/cmdline"
	"golite/lexer"
	"golite/parser"
)

func main() {
	// Parse the command line
	cmdline := cmdline.ReadCmdline()
	if cmdline == nil {
		return
	}

	// Get the input source file path
	// Input source path is always the last argument
	inputSourcePath := cmdline.Args[len(cmdline.Args)-1]

	// Get a new lexer
	lexer := lexer.NewLexer(inputSourcePath)

	// Get a new parser
	parser := parser.NewParser(lexer)

	// Parse the input source file (lexes and parses)
	invalidParse := parser.Parse()
	if invalidParse {
		fmt.Println("Valid parse as program is correct (till now...)!")
	} else {
		fmt.Println("Invalid parse as program is incorrect!")
	}
}
