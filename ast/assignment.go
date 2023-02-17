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
func (a *Assignment) BuildSymbolTable(tables *st.SymbolTables, errors []*SemanticAnalysisError) []*SemanticAnalysisError {
	// Nothing to do here since an assignment is not added to the symbol table
	return errors
}

// Type checking for the assignment
func (a *Assignment) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables) []*SemanticAnalysisError {
	// Type check the expression on the right hand side of the assignment
	// Also ensures that GetType() on expression is the correct predictive type
	errors = a.right.TypeCheck(errors, tables)

	// Type check the lvalue on the left hand side of the assignment
	errors = a.variable.TypeCheck(errors, tables)

	// Check that the type of the lvalue is the same as the type of the expression
	if a.variable.GetType() != a.right.GetType() {
		errors = append(errors, NewSemanticAnalysisError("Type mismatch in assignment", "type mistmatch", a.GetToken()))
	}

	return errors
}

// Get the token for the assignment
func (a *Assignment) GetToken() *token.Token {
	return a.Token
}
