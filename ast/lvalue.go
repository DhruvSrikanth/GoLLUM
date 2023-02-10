package ast

import (
	"bytes"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

// LValue node in the AST
type LValue struct {
	*token.Token
	identifier string     // The identifier of the variable
	fields     []string   // The field of the struct
	ty         types.Type // The type of the lvalue
}

// NewLValue node
func NewLValue(tok *token.Token, identifier string, fields []string, ty types.Type) *LValue {
	return &LValue{tok, identifier, fields, ty}
}

// String representation of the lvalue node
func (l *LValue) String() string {
	var out bytes.Buffer

	out.WriteString(l.identifier)
	// Its possible that the lvalue is an struct field
	for _, field := range l.fields {
		out.WriteString(".")
		out.WriteString(field)
	}

	return out.String()
}

// Build the symbol table for the lvalue node
func (l *LValue) BuildSymbolTable(tables *st.SymbolTables) {
	// Nothing to do here since a lvalue is not added to the symbol table since it is already declared
}

// Type check the lvalue node
func (l *LValue) TypeCheck(errors []string, tables *st.SymbolTables) []string {
	return errors
}

// Get the type of the lvalue node
func (l *LValue) GetType() types.Type {
	return l.ty
}
