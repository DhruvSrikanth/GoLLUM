package main

import (
	"fmt"
	"golite/cmdline"
	"golite/lexer"
	"golite/parser"
	sa "golite/semanticanalysis"
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

	// Check the lex flag
	if cmd.LexFlag {
		// Print the tokens
		lexer.PrintTokens()
	} else {
		// Get a new parser
		parser := parser.NewParser(lexer)

		// Parse the input source file (lexes and parses)
		// and get the AST
		if ast := parser.Parse(); ast != nil {
			if cmd.ASTFlag {
				// Print the AST
				fmt.Println(ast)
			}
			if tables := sa.Execute(ast); tables != nil {
				fmt.Println("Passed semantic analysis (view console for errors at lexing and parsing stages)")
			} else {
				fmt.Println("Failed semantic analysis (view console for errors at lexing, parsing and semantic analysis stages)")
			}
		}
	}
}
