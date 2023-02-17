package ast

import (
	"bytes"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

// Read node struct for the AST
type Read struct {
	*token.Token
	lval LValue
}

// New Read node
func NewRead(lval LValue, token *token.Token) *Read {
	return &Read{token, lval}
}

// String representation of the read node
func (r *Read) String() string {
	var out bytes.Buffer

	out.WriteString("scan ")
	out.WriteString(r.lval.String())
	out.WriteString(";")

	return out.String()
}

// Build the symbol table for the read node
func (r *Read) BuildSymbolTable(tables *st.SymbolTables, errors []*SemanticAnalysisError) []*SemanticAnalysisError {
	// Nothing to do here since a read statement uses a variable already declared in the symbol table
	return errors
}

// Type check the read node*SemanticAnalysisError
func (r *Read) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables, funcEntry *st.FuncEntry) []*SemanticAnalysisError {
	// Type check the lvalue
	errors = r.lval.TypeCheck(errors, tables, funcEntry)

	// Get the type of the lvalue
	lvalType := r.lval.GetType()

	// Check if the lvalue is a valid type (only int is valid)
	if lvalType != types.StringToType("int") {
		// Set to the valid type for predictive type checking
		r.lval.ty = types.StringToType("int")
		errors = append(errors, NewSemanticAnalysisError("Read statement can only be used with int variables", "type mismatch", r.Token))
	}

	return errors
}
