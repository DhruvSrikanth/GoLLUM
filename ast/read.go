package ast

import (
	"bytes"
	st "golite/symboltable"
	"golite/token"
)

// Read node struct for the AST
type Read struct {
	*token.Token
	lval LValue
}

// New Read node
func NewRead(lval LValue, token *token.Token) *Read {
	return &Read{token, lval}
}

// String representation of the read node
func (r *Read) String() string {
	var out bytes.Buffer

	out.WriteString("scan ")
	out.WriteString(r.lval.String())
	out.WriteString(";")

	return out.String()
}

// Build the symbol table for the read node
func (r *Read) BuildSymbolTable(tables *st.SymbolTables) {
	// Nothing to do here since a read statement uses a variable already declared in the symbol table
}

// Type check the read node
func (r *Read) TypeCheck(errors []string, tables *st.SymbolTables) []string {
	return errors
}
