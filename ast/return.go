package ast

import (
	"bytes"
	st "golite/symboltable"
	"golite/token"
)

// Return production rule node in the AST
type Return struct {
	*token.Token
	expression *Expression // The expression to be returned
}

// NewReturn node
func NewReturn(expression *Expression, token *token.Token) *Return {
	return &Return{token, expression}
}

// String representation of the return node
func (r *Return) String() string {
	var out bytes.Buffer

	out.WriteString("return ")
	if r.expression != nil {
		out.WriteString((*r.expression).String())
	}
	out.WriteString(";")

	return out.String()
}

// Build the symbol table for the return node
func (r *Return) BuildSymbolTable(tables *st.SymbolTables) {
	// Nothing to do here since the expression variables are already in the symbol table
}

// Type checking for the return node
func (r *Return) TypeCheck(errors []string, tables *st.SymbolTables) []string {
	return errors
}
