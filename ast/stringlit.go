package ast

import (
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

// String literal
type StringLiteral struct {
	*token.Token            // The token for the string literal
	Value        string     // The value for the string literal
	ty           types.Type // The type of the literal
}

// New instance of the string literal
func NewStringLit(token *token.Token, value string) *StringLiteral {
	return &StringLiteral{token, value, types.StringToType("string")}
}

// String representation of the string literal
func (sl *StringLiteral) String() string {
	return sl.Value
}

// Get the inferred type of the literal
func (sl *StringLiteral) GetType() types.Type {
	return sl.ty
}

// TypeCheck performs type checking for the string literal
func (sl *StringLiteral) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables, funcEntry *st.FuncEntry) []*SemanticAnalysisError {
	return errors
}
