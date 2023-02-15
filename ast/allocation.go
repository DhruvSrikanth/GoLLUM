package ast

import (
	"bytes"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

// Allocate node in the AST
type Allocate struct {
	*token.Token
	structType string
	ty         types.Type
}

// New Delete node
func NewAllocate(structType string, token *token.Token) *Allocate {
	return &Allocate{token, structType, nil}
}

// String representation of the allocate node
func (d *Allocate) String() string {
	var out bytes.Buffer

	out.WriteString("new ")
	out.WriteString(d.structType)

	return out.String()
}

// Build the symbol table for the allocate node
func (d *Allocate) BuildSymbolTable(tables *st.SymbolTables) {
	// Nothing to do here since allocate is on a variable thats already been declared and added to the symbol table
}

// Type check the allocate node
func (d *Allocate) TypeCheck(errors []string, tables *st.SymbolTables) []string {
	return errors
}

// Get the type of the allocate node
func (d *Allocate) GetType() types.Type {
	return d.ty
}
