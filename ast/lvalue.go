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
func NewLValue(identifier string, fields []string, ty types.Type, tok *token.Token) *LValue {
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
func (l *LValue) BuildSymbolTable(tables *st.SymbolTables, errors []*SemanticAnalysisError) []*SemanticAnalysisError {
	// Nothing to do here since a lvalue is not added to the symbol table since it is already declared
	return errors
}

// Type check the lvalue node
func (l *LValue) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables) []*SemanticAnalysisError {
	// Check if the identifer is declared
	entry := tables.Funcs.Contains(l.identifier)
	if entry == nil {
		errors = append(errors, NewSemanticAnalysisError("Undeclared variable "+l.identifier+".", "undeclared variable", l.Token))
	} else {
		if len(l.fields) == 0 {
			// The lvalue is a variable
			l.ty = entry.GetType()
		} else {
			for _, field := range l.fields {
				entryType := entry.GetType()
				if types.TypeToKind(entryType) == types.STRUCT {
					entryStructName := entryType.String()[1:] // Remove the * from the type name
					// Check if the struct exists in the symbol table
					structEntry := tables.Structs.Contains(entryStructName)
					if structEntry == nil {
						errors = append(errors, NewSemanticAnalysisError("Type "+entryType.String()+" not declared.", "undeclared type", l.Token))
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
							errors = append(errors, NewSemanticAnalysisError("Field "+field+" not declared in type "+entryType.String(), "undeclared field", l.Token))
							break
						}
						// The field exists in the struct so move on to the next field
						entryType = matchedField
					}
				} else {
					// Primitive type meaning it must be the last field
					if field != l.fields[len(l.fields)-1] {
						errors = append(errors, NewSemanticAnalysisError("Cannot access field "+field+" of type "+entryType.String(), "field access error", l.Token))
					}

					// Set the type of the lvalue
					l.ty = entryType
				}
			}
		}
	}

	return errors
}

// Get the type of the lvalue node
func (l *LValue) GetType() types.Type {
	return l.ty
}
