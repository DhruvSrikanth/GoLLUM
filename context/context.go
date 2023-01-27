package context

import (
	"fmt"
)

// Alias for which phase the compiler is in
type CompilerPhase int

// Define the different compiler phases here
const (
	LEXER CompilerPhase = iota
	PARSER
)

// CompilerError struct (how to handle errors)
type CompilerError struct {
	Line   int           // The line where the error occurred
	Column int           // The column where the error occurred
	Msg    string        // A user friendly message to display
	Phase  CompilerPhase // The phase of the compiler were the error was generated
}

// to_str() method for CompilerError i.e. returning a beautified string
func (err *CompilerError) String() string {
	if err.Phase == LEXER {
		return fmt.Sprintf("lexer error(%d:%d): %s", err.Line, err.Column, err.Msg)
	} else if err.Phase == PARSER {
		return fmt.Sprintf("syntax error(%d,%d): %s", err.Line, err.Column, err.Msg)
	}
	panic("Invalid phase found")
}

// Check if the compiler has errors
func HasErrors(errs []*CompilerError) bool {
	if len(errs) > 0 {
		for _, e := range errs {
			fmt.Println(e)
		}
		return true
	}
	return false
}
