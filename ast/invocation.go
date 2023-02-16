package ast

import (
	"bytes"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

// Invocation node for the AST
type Invocation struct {
	*token.Token
	identifier string       // The identifier of the function being invoked
	arguments  []Expression // The arguments to the function
	ty         types.Type
}

// NewInvocation node
func NewInvocation(identifier string, arguments []Expression, token *token.Token) *Invocation {
	return &Invocation{token, identifier, arguments, nil}
}

// String representation of the invocation node
func (i *Invocation) String() string {
	var out bytes.Buffer

	out.WriteString(i.identifier)
	out.WriteString("(")
	for x, arg := range i.arguments {
		out.WriteString(arg.String())
		if x < len(i.arguments)-1 {
			out.WriteString(", ")
		}
	}
	out.WriteString(")")
	out.WriteString(";")

	return out.String()
}

// Build the symbol table for the invocation node
func (i *Invocation) BuildSymbolTable(tables *st.SymbolTables, errors []*SemanticAnalysisError) []*SemanticAnalysisError {
	// Nothing to do here since the function is already in the symbol table and the invocation is not a declaration
	return errors
}

// Type checking for the invocation node
func (i *Invocation) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables) []*SemanticAnalysisError {
	return errors
}

// Get the type of the invocation node
func (i *Invocation) GetType() types.Type {
	return i.ty
}
