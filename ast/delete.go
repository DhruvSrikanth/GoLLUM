package ast

import (
	"bytes"
	st "golite/symboltable"
	"golite/token"
)

// Delete node in the AST
type Delete struct {
	*token.Token
	lval LValue
}

// New Delete node
func NewDelete(lval LValue, token *token.Token) *Delete {
	return &Delete{token, lval}
}

// String representation of the delete node
func (d *Delete) String() string {
	var out bytes.Buffer

	out.WriteString("delete ")
	out.WriteString(d.lval.String())
	out.WriteString(";")

	return out.String()
}

// Build the symbol table for the delete node
func (d *Delete) BuildSymbolTable(tables *st.SymbolTables) {
	// Nothing to do here since a delete is not added to the symbol table
}

// Type check the delete node
func (d *Delete) TypeCheck(errors []string, tables *st.SymbolTables) []string {
	return errors
}
