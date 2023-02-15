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
func (v *VariableInvocation) BuildSymbolTable(tables *st.SymbolTables) {
	// Nothing to do here since the function is already in the symbol table and the VariableInvocation is not a declaration
}

// Type checking for the VariableInvocation node
func (v *VariableInvocation) TypeCheck(errors []string, tables *st.SymbolTables) []string {
	return errors
}

// Get the type of the VariableInvocation node
func (v *VariableInvocation) GetType() types.Type {
	return v.ty
}
