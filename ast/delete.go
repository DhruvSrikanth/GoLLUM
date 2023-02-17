package ast

import (
	"bytes"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

// Delete node in the AST
type Delete struct {
	*token.Token
	expr Expression // Expression to be deleted
}

// New Delete node
func NewDelete(expr Expression, token *token.Token) *Delete {
	return &Delete{token, expr}
}

// String representation of the delete node
func (d *Delete) String() string {
	var out bytes.Buffer

	out.WriteString("delete ")
	out.WriteString(d.expr.String())
	out.WriteString(";")

	return out.String()
}

// Build the symbol table for the delete node
func (d *Delete) BuildSymbolTable(tables *st.SymbolTables, errors []*SemanticAnalysisError) []*SemanticAnalysisError {
	// Nothing to do here since a delete is not added to the symbol table
	return errors
}

// Type check the delete node
func (d *Delete) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables, funcEntry *st.FuncEntry) []*SemanticAnalysisError {
	// Type check the expression
	errors = d.expr.TypeCheck(errors, tables, funcEntry)

	// Check if expression has been declared
	if tables.Funcs.Contains(d.expr.String()) == nil {
		errors = append(errors, NewSemanticAnalysisError(d.expr.String()+" has not been declared.", "undeclared identifier", d.Token))
	}

	// Check if the expression is a pointer to a struct type
	exprType := d.expr.GetType()
	if types.TypeToKind(exprType) != types.STRUCT {
		errors = append(errors, NewSemanticAnalysisError("delete can only be called on pointers to non-primitive types.", "invalid memory management", d.Token))
	} else if tables.Structs.Contains(exprType.String()[1:]) == nil {
		// Check if struct declared in symbol table
		errors = append(errors, NewSemanticAnalysisError("struct "+exprType.String()[1:]+" does not exist.", "invalid type", d.Token))
	}
	return errors
}
