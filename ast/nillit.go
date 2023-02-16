package ast

import (
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

// The nil literal
type NilLiteral struct {
	*token.Token            // The token for the integer literal
	Value        string     // The value for the integer literal
	ty           types.Type // The type of the literal
}

// Create a new nil literal
func NewNilLit(token *token.Token) *NilLiteral {
	return &NilLiteral{token, "nil", types.StringToType("nil")}
}

// String representation of the nil literal
func (nl *NilLiteral) String() string {
	return nl.Value
}

// Get the inferred type of the literal
func (nl *NilLiteral) GetType() types.Type {
	return nl.ty
}

// TypeCheck performs type checking for the nil literal
func (nl *NilLiteral) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables) []*SemanticAnalysisError {
	return errors
}
