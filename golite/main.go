package main

import (
	"fmt"
	"golite/cmdline"
	"golite/lexer"
	"golite/llvm"
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
	targetMachine := "arm64-apple-darwin22.2.0"

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
			} else {
				if tables := sa.Execute(ast); tables != nil {
					// Translate the all table types to their LLVM representation
					tables.LLVMTypeConversion()
					if cmd.SymbolTableFlag {
						// Print the symbol table
						fmt.Println(tables)
					}

					// Continue translation from AST to LLVM IR
					// Create the target information
					targetInfo := llvm.NewTargetInformation(inputSourcePath, targetMachine)
					fmt.Println(targetInfo)

				} else {
					fmt.Println("Failed semantic analysis")
				}
			}
		}
	}
}
