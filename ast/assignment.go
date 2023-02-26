package ast

import (
	"bytes"
	"golite/llvm"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

// Assignment node struct for the AST

type Assignment struct {
	*token.Token
	variable LValue     // The lvalue for the assignment statement
	right    Expression // The value assigned to the variable of this statements
}

// New Assignment node
func NewAssignment(variable LValue, right Expression, token *token.Token) *Assignment {
	return &Assignment{token, variable, right}
}

// String representation of the assignment node
func (a *Assignment) String() string {
	var out bytes.Buffer

	out.WriteString(a.variable.String())
	out.WriteString(" = ")
	out.WriteString(a.right.String())
	out.WriteString(";")

	return out.String()
}

// Build the symbol table for the assignment
func (a *Assignment) BuildSymbolTable(tables *st.SymbolTables, errors []*SemanticAnalysisError) []*SemanticAnalysisError {
	// Nothing to do here since an assignment is not added to the symbol table
	return errors
}

// Type checking for the assignment
func (a *Assignment) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables, funcEntry *st.FuncEntry) []*SemanticAnalysisError {
	// Type check the expression on the right hand side of the assignment
	// Also ensures that GetType() on expression is the correct predictive type
	errors = a.right.TypeCheck(errors, tables, funcEntry)

	// Type check the lvalue on the left hand side of the assignment
	errors = a.variable.TypeCheck(errors, tables, funcEntry)

	if a.variable.GetType() == nil {
		// Use inferred data type from rhs to continue using predictive type checking
		// This ensures that the compiler can continue type checking to find more errors
		// Error for this nil type will be added in the type checking of the lvalue so all errors can be accounted for
		a.variable.ty = a.right.GetType()
		errors = append(errors, NewSemanticAnalysisError("type mismatch in assignment", "type mistmatch", a.GetToken()))
	}

	// Check that the type of the lvalue is the same as the type of the expression
	// Check if it is a struct type
	if a.variable.GetType() != a.right.GetType() {
		// If the types are not the same, check if the lvalue is a struct type
		// If it is, the right hand side must also be a struct literal or nil
		// In this case since it is not a struct literal or atleast not the same struct literal, we can check for nil
		if (types.TypeToKind(a.variable.GetType()) == types.STRUCT) && (a.right.GetType() == types.StringToType("nil")) {
			return errors
		}
		errors = append(errors, NewSemanticAnalysisError("type mismatch in assignment", "type mistmatch", a.GetToken()))
	}

	return errors
}

// Get the token for the assignment
func (a *Assignment) GetToken() *token.Token {
	return a.Token
}

// Control flow analysis for the assignment node
func (a *Assignment) GetControlFlow(errors []*SemanticAnalysisError, funcEntry *st.FuncEntry) ([]*SemanticAnalysisError, bool) {
	// Nothing to do here since the assignment node does not have any control flow
	return errors, false
}

// Translate the assignment node to LLVM IR
func (a *Assignment) ToLLVMCFG(tables *st.SymbolTables, blocks []*llvm.BasicBlock, funcEntry *st.FuncEntry) []*llvm.BasicBlock {
	return blocks
}
