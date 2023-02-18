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
func (s *SelectorTerm) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables, funcEntry *st.FuncEntry) []*SemanticAnalysisError {
	// Type check the factor
	errors = s.factor.TypeCheck(errors, tables, funcEntry)

	// Check if the identifer is declared
	identifer := s.factor.String()
	entry := funcEntry.Variables.Contains(identifer)
	if entry == nil {
		// Check the parameters
		for _, param := range funcEntry.Parameters {
			if param.Name == identifer {
				entry = param
				break
			}
		}
	}
	if entry == nil {
		// Check if the identifier is a literal
		if s.factor.GetType() == types.StringToType("int") || s.factor.GetType() == types.StringToType("bool") || s.factor.GetType() == types.StringToType("nil") {
			s.ty = s.factor.GetType()
			return errors
		} else {
			errors = append(errors, NewSemanticAnalysisError("undeclared variable "+identifer, "undeclared variable", s.Token))
		}
	} else {
		if len(s.fields) == 0 {
			// The lvalue is a variable
			s.ty = entry.GetType()
		} else {
			for _, field := range s.fields {
				entryType := entry.GetType()
				if types.TypeToKind(entryType) == types.STRUCT {
					entryStructName := entryType.String()[1:] // Remove the * from the type name
					// Check if the struct exists in the symbol table
					structEntry := tables.Structs.Contains(entryStructName)
					if structEntry == nil {
						errors = append(errors, NewSemanticAnalysisError("type "+entryType.String()+" not declared.", "undeclared type", s.Token))
					} else {
						// Check if the field exists in the struct
						found := false
						var matchedField *st.FieldEntry
						for _, fieldEntry := range structEntry.Fields {
							if fieldEntry.Name == field {
								found = true
								matchedField = fieldEntry
								break
							}
						}
						if !found {
							errors = append(errors, NewSemanticAnalysisError("field "+field+" not declared in type "+entryType.String(), "undeclared field", s.Token))

							// Set the type to nil
							s.ty = types.StringToType("nil")
							break
						}
						// The field exists in the struct so move on to the next field
						entryType = matchedField
					}
				} else {
					// Primitive type meaning it must be the last field
					if field != s.fields[len(s.fields)-1] {
						errors = append(errors, NewSemanticAnalysisError("cannot access field "+field+" of type "+entryType.String(), "field access error", s.Token))
					}

					// Set the type of the lvalue
					s.ty = entryType
				}
			}
		}
	}

	return errors
}

// Get the type of the SelectorTerm node
func (s *SelectorTerm) GetType() types.Type {
	return s.ty
}
