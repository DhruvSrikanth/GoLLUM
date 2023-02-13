package ast

import (
	"fmt"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

// Boolean literal
type BoolLiteral struct {
	*token.Token            // The token for the integer literal
	Value        bool       // The value for the integer literal
	ty           types.Type // The type of the literal
}

// Creates a new instance of a booleal literal
func NewBoolLit(token *token.Token, value bool) *BoolLiteral {
	return &BoolLiteral{token, value, types.StringToType("bool")}
}

// String representation of the bool literal
func (bl *BoolLiteral) String() string {
	return fmt.Sprintf("%d", bl.Value)
}

// Get the inferred type of the literal
func (bl *BoolLiteral) GetType() types.Type {
	return bl.ty
}

// TypeCheck performs type checking for the nil literal
func (bl *BoolLiteral) TypeCheck(errors []string, tables *st.SymbolTables) []string {
	return errors
}