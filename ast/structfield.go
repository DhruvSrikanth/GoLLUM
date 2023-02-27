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

	out.WriteString(d.variable)
	out.WriteString(" ")
	out.WriteString(d.ty.String())
	return out.String()
}

// Build the symbol table for the declaration
func (d *Decl) BuildSymbolTable(tables *st.SymbolTables, errors []*SemanticAnalysisError) []*SemanticAnalysisError {
	// Nothing to do since we add the decls to the symbol table in the struct decl BuildSymbolTable function
	return errors
}

func (d *Decl) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables, funcEntry *st.FuncEntry) []*SemanticAnalysisError {
	//When you reach a declaration you want to make sure you add the variable declaration to the symbol table.
	return errors
}

// Get the type of the declaration
func (d *Decl) GetType() types.Type {
	return d.ty
}

// Get the name of the variable declared
func (d *Decl) GetVariable() string {
	return d.variable
}
