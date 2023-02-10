package ast

import (
	"fmt"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

// Integer literal
type IntLiteral struct {
	*token.Token            // The token for the integer literal
	Value        int64      // The value for the integer literal
	ty           types.Type // The type of the literal
}

// Creates a new instance of an integer literal
func NewIntLit(value int64, token *token.Token) *IntLiteral {
	return &IntLiteral{token, value, types.StringToType("int")}
}

// String representation of the integer literal
func (il *IntLiteral) String() string {
	return fmt.Sprintf("%d", il.Value)
}

// Get the inferred type of the string literal
func (il *IntLiteral) GetType() types.Type {
	return il.ty
}

// TypeCheck performs type checking for the integer literal
func (il *IntLiteral) TypeCheck(errors []string, tables *st.SymbolTables) []string {
	return errors
}
