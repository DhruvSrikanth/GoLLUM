package main

import (
	"golite/cmdline"
	"golite/lexer"
	"golite/parser"
)

func main() {
	// Parse the command line
	cmdline := cmdline.ReadCmdline()

	// Get the input source file path
	inputSourcePath := cmdline.Args[0]

	// Get a new lexer
	lexer := lexer.NewLexer(inputSourcePath)

	// Get a new parser
	parser := parser.NewParser(lexer)

	// Parse the input source file (lexes and parses)
	parser.Parse()
}
