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
	ty     types.Type // The type of the selector term
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
	// Check if the selector is a struct field or a variable/function
	// if len(s.fields) > 0 {
	// 	// Struct field
	// 	// Get the variable name
	// 	varName := s.factor.(*VariableInvocation).identifier
	// 	// Get the variable entry
	// 	entry := tables.Funcs.Contains(varName)
	// 	// Check if the variable exists
	// 	if entry == nil {
	// 		errors = append(errors, NewSemanticAnalysisError("Variable "+varName+" not declared.", "undeclared", s.Token))
	// 	} else {
	// 		// Check if the variable is a struct thats been declared
	// 		structType := entry.GetType()
	// 		structEntry := tables.Structs.Contains(structType)
	// 		if entry.GetType() == nil {
	// 			errors = append(errors, NewSemanticAnalysisError("Variable "+varName+" is not a struct.", "not a struct", s.Token))
	// 		}
	// 	}
	// }
	return errors
}

// Get the type of the SelectorTerm node
func (s *SelectorTerm) GetType() types.Type {
	return s.ty
}
