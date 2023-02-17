package ast

import (
	"bytes"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

// VariableInvocation node for the AST
type VariableInvocation struct {
	*token.Token
	identifier string       // The identifier of the variable being invoked
	arguments  []Expression // The arguments to the variable being invoked (function call case)
	ty         types.Type   // The type of the variable being invoked (or return type of the function being called)
}

// NewVariableInvocation node
func NewVariableInvocation(identifier string, arguments []Expression, token *token.Token) *VariableInvocation {
	return &VariableInvocation{token, identifier, arguments, nil}
}

// String representation of the variable invocation node
func (v *VariableInvocation) String() string {
	var out bytes.Buffer

	out.WriteString(v.identifier)
	if len(v.arguments) > 0 {
		out.WriteString("(")
		for x, arg := range v.arguments {
			out.WriteString(arg.String())
			if x < len(v.arguments)-1 {
				out.WriteString(", ")
			}
		}
		out.WriteString(")")
	}

	return out.String()
}

// Build the symbol table for the VariableInvocation node
func (v *VariableInvocation) BuildSymbolTable(tables *st.SymbolTables, errors []*SemanticAnalysisError) []*SemanticAnalysisError {
	// Nothing to do here since the function is already in the symbol table and the VariableInvocation is not a declaration
	return errors
}

// Type checking for the VariableInvocation node
func (v *VariableInvocation) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables) []*SemanticAnalysisError {
	// Check if its a variable or function call
	entry := tables.Funcs.Contains(v.identifier)
	if len(v.arguments) != 0 {
		// Function call
		// Check if the function exists in the symbol table
		if entry == nil {
			errors = append(errors, NewSemanticAnalysisError("Function "+v.identifier+" not declared.", "undeclared", v.Token))
		} else {
			// Check if the number of arguments is correct\
			if len(entry.Parameters) != len(v.arguments) {
				errors = append(errors, NewSemanticAnalysisError("Function "+v.identifier+" expects "+string(len(entry.Parameters))+" arguments but "+string(len(v.arguments))+" were provided.", "invalid number of arguments", v.Token))
			} else {
				// Call type check on each argument
				// to ensure the GetTypes() method returns the correct type even if there are errors
				for _, arg := range v.arguments {
					errors = arg.TypeCheck(errors, tables)
				}
				// Check if the types of the arguments are correct
				for i, arg := range v.arguments {
					if arg.GetType() != entry.Parameters[i].GetType() {
						errors = append(errors, NewSemanticAnalysisError("Expected type "+entry.Parameters[i].GetType().String()+" but got "+arg.GetType().String()+".", "invalid argument type", v.Token))
					}
				}
			}
		}

	} else {
		// Variable
		// Check if the variable exists in the symbol table
		if entry == nil {
			errors = append(errors, NewSemanticAnalysisError("Variable "+v.identifier+" not declared.", "undeclared", v.Token))
		}
	}

	// Set the type of the variable invocation
	v.ty = entry.GetType()
	return errors
}

// Get the type of the VariableInvocation node
func (v *VariableInvocation) GetType() types.Type {
	return v.ty
}
