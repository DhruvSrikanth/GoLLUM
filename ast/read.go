package ast

import (
	"bytes"
	"golite/llvm"
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
		errors = append(errors, NewSemanticAnalysisError("read statement can only be used with int variables", "type mismatch", r.Token))
	}

	return errors
}

// Control flow analysis for the read node
func (r *Read) GetControlFlow(errors []*SemanticAnalysisError, funcEntry *st.FuncEntry) ([]*SemanticAnalysisError, bool) {
	// Nothing to do here since the read node does not have any control flow
	return errors, false
}

// Translate the read node to LLVM IR
func (r *Read) ToLLVMCFG(tables *st.SymbolTables, blocks []*llvm.BasicBlock, funcEntry *st.FuncEntry) []*llvm.BasicBlock {
	// Stay in the same block
	return blocks
}
