package ast

import (
	"bytes"
	st "golite/symboltable"
	"golite/token"
)

// Assignment node struct for the AST

type Assignment struct {
	*token.Token
	variable LValue     // The lvalue for the assignment statement
	right    Expression // The value assigned to the variable of this statements
	// The right value could either be an expression or a scan statement
}

// New Assignment node
func NewAssignment(variable LValue, right Expression, token *token.Token) *Assignment {
	return &Assignment{token, variable, right}
}

// String representation of the assignment node
func (a *Assignment) String() string {
	var out bytes.Buffer

	out.WriteString(a.variable.String())
	out.WriteString(" = ")
	out.WriteString(a.right.String())
	out.WriteString(";")

	return out.String()
}

// Build the symbol table for the assignment
func (a *Assignment) BuildSymbolTable(tables *st.SymbolTables) {
	// Nothing to do here since an assignment is not added to the symbol table
}

// Type checking for the assignment
func (a *Assignment) TypeCheck(errors []string, tables *st.SymbolTables) []string {
	return errors
}
