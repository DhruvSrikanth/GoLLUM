package ast

import (
	"bytes"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

// UnaryTerm node
type UnaryTerm struct {
	*token.Token
	unaryTrm SelectorTerm // The identifier of the variable
	op       *Operator    // The operator for the unary expression
	ty       types.Type   // The type of the lvalue
}

// NewUnaryTerm node
func NewUnaryTerm(unaryTrm SelectorTerm, op *Operator, ty types.Type, tok *token.Token) *UnaryTerm {
	return &UnaryTerm{tok, unaryTrm, op, ty}
}

// String representation of the lvalue node
func (s *UnaryTerm) String() string {
	var out bytes.Buffer
	if s.op != nil {
		out.WriteString(OpToStr(*s.op))
	}

	out.WriteString(s.unaryTrm.String())

	return out.String()
}

// Build the symbol table for the UnaryTerm node
func (u *UnaryTerm) BuildSymbolTable(tables *st.SymbolTables) {
	// Nothing to do here since a lvalue is not added to the symbol table since it is already declared
}

// Type check the UnaryTerm node
func (u *UnaryTerm) TypeCheck(errors []string, tables *st.SymbolTables) []string {
	return errors
}

// Get the type of the UnaryTerm node
func (u *UnaryTerm) GetType() types.Type {
	return u.ty
}
