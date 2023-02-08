package ast

import (
	"bytes"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

// Decl is a declaration production rule which for a struct field
type Decl struct {
	*token.Token
	variable string     // lvalue (no rvalue since only lvalues are allowed in declarations)
	ty       types.Type // type of variable
}

// Returns an instance to a new declaration
func NewDeclStm(variable string, ty types.Type, token *token.Token) *Decl {
	return &Decl{token, variable, ty}
}

// String representation of the declaration
func (d *Decl) String() string {
	var out bytes.Buffer

	out.WriteString(d.ty.String())
	out.WriteString(" ")
	out.WriteString(d.variable)
	return out.String()
}

// Build the symbol table for the declaration
func (d *Decl) BuildSymbolTable(tables *st.SymbolTables) {
	// Nothing to do since we add the decls to the symbol table in the struct decl BuildSymbolTable function
}

func (d *Decl) TypeCheck(errors []string, tables *st.SymbolTables) []string {
	//When you reach a declaration you want to make sure you add the variable declaration to the symbol table.
	return errors
}