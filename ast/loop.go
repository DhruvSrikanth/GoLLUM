package ast

import (
	"bytes"
	st "golite/symboltable"
	"golite/token"
)

// Loop node for the AST
type Loop struct {
	*token.Token
	condition Expression // The condition for the loop
	body      Block      // The block to be executed if the condition is true
}

// NewLoop node
func NewLoop(condition Expression, body Block, token *token.Token) *Loop {
	return &Loop{token, condition, body}
}

// String representation of the loop node
func (l *Loop) String() string {
	var out bytes.Buffer

	out.WriteString("for (")
	out.WriteString(l.condition.String())
	out.WriteString(") ")
	out.WriteString(l.body.String())

	return out.String()
}

// Build the symbol table for the loop node
func (l *Loop) BuildSymbolTable(tables *st.SymbolTables, errors []*SemanticAnalysisError) []*SemanticAnalysisError {
	// Nothing to do here since the body is already in the symbol table
	return errors
}

// Type checking for the loop node
func (l *Loop) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables) []*SemanticAnalysisError {
	return errors
}
