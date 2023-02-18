package ast

import (
	"bytes"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

// Return production rule node in the AST
type Return struct {
	*token.Token
	expression *Expression // The expression to be returned
	ty         types.Type  // The type of the expression
}

// NewReturn node
func NewReturn(expression *Expression, token *token.Token) *Return {
	return &Return{token, expression, nil}
}

// String representation of the return node
func (r *Return) String() string {
	var out bytes.Buffer

	out.WriteString("return ")
	if *r.expression != nil {
		out.WriteString((*r.expression).String())
	}
	out.WriteString(";")

	return out.String()
}

// Build the symbol table for the return node
func (r *Return) BuildSymbolTable(tables *st.SymbolTables, errors []*SemanticAnalysisError) []*SemanticAnalysisError {
	// Nothing to do here since the expression variables are already in the symbol table
	return errors
}

// Type checking for the return node
func (r *Return) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables, funcEntry *st.FuncEntry) []*SemanticAnalysisError {
	// If the expression does not exist, then the return type is void
	if *r.expression == nil {
		r.ty = types.StringToType("void")
	} else {
		// Type check the expression
		// Ensures that GetType() has a valid type
		errors = (*r.expression).TypeCheck(errors, tables, funcEntry)

		// Get the type of the expression
		r.ty = (*r.expression).GetType()
	}

	// Check if the return type is the same as the function return type
	if r.ty != funcEntry.RetTy {
		errors = append(errors, NewSemanticAnalysisError("return type does not match function signature return type", "mismatched types", r.Token))
	}

	return errors
}

// Get the type of the return node
func (r *Return) GetType() types.Type {
	return r.ty
}
