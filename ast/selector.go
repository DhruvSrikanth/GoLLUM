package ast

import (
	"bytes"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

type SelectorTerm struct {
	*token.Token
	factor Expression // The identifier of the variable
	fields []string   // The field of the struct (if it is a struct)
	ty     types.Type // The type of the unary term
}

// NewSelectorTerm node
func NewSelectorTerm(factor Expression, fields []string, ty types.Type, tok *token.Token) *SelectorTerm {
	return &SelectorTerm{tok, factor, fields, ty}
}

// String representation of the SelectorTerm node
func (s *SelectorTerm) String() string {
	var out bytes.Buffer

	out.WriteString(s.factor.String())
	// Its possible that the SelectorTerm is an struct field
	for _, field := range s.fields {
		out.WriteString(".")
		out.WriteString(field)
	}

	return out.String()
}

// Build the symbol table for the SelectorTerm node
func (s *SelectorTerm) BuildSymbolTable(tables *st.SymbolTables, errors []*SemanticAnalysisError) []*SemanticAnalysisError {
	// Nothing to do here since a SelectorTerm is not added to the symbol table since it is already declared
	return errors
}

// Type check the SelectorTerm node
func (s *SelectorTerm) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables) []*SemanticAnalysisError {
	return errors
}

// Get the type of the SelectorTerm node
func (s *SelectorTerm) GetType() types.Type {
	return s.ty
}
