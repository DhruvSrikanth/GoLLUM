package ast

import (
	"bytes"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

// Struct for declaration production rule
type Declaration struct {
	*token.Token
	variable string     // Variable name
	ty       types.Type // Type of the variable
}

// Create a new declaration node
func NewDeclaration(variable string, ty types.Type, token *token.Token) *Declaration {
	return &Declaration{token, variable, ty}
}

// String representation of the declaration node
func (d *Declaration) String() string {
	var out bytes.Buffer
	out.WriteString("var ")
	out.WriteString(d.variable)
	out.WriteString(" ")
	out.WriteString(d.ty.String())
	out.WriteString(";")
	return out.String()
}

// Build the symbol table for the declaration
func (d *Declaration) BuildSymbolTable(tables *st.SymbolTables) {
	tables.Globals.Insert(d.variable, &st.VarEntry{d.variable, d.ty, st.GLOBAL})
}

// Type checking for the declaration
func (d *Declaration) TypeCheck(errors []string, tables *st.SymbolTables) []string {
	return errors
}
