package llvm

import (
	"bytes"
)

// Representation of a function declaration
type FunctionDecl struct {
	// The name of the function declaration
	name   string
	fields string
}

// NewStructDecl returns a new function declaration
func NewFunctionDecl() *FunctionDecl {
	return &FunctionDecl{"name", "fields"}
}

// String representation of the function declaration
func (f *FunctionDecl) String() string {
	var out bytes.Buffer
	return out.String()
}
