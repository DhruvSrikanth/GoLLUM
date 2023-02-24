package ast

import (
	"bytes"
	"golite/llvm"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

// Struct for declaration production rule
type Declaration struct {
	*token.Token
	variable string     // Variable name
	ty       types.Type // Type of the variable
}

// Create a new declaration node
func NewDeclaration(variable string, ty types.Type, token *token.Token) *Declaration {
	return &Declaration{token, variable, ty}
}

// String representation of the declaration node
func (d *Declaration) String() string {
	var out bytes.Buffer
	out.WriteString("var ")
	out.WriteString(d.variable)
	out.WriteString(" ")
	out.WriteString(d.ty.String())
	out.WriteString(";")
	return out.String()
}

// Build the symbol table for the declaration
func (d *Declaration) BuildSymbolTable(tables *st.SymbolTables, errors []*SemanticAnalysisError) []*SemanticAnalysisError {
	// No need to add things to the symbol table here because declarations are added to the symbol table in the program
	// for global declarations and in the function for local declarations
	return errors
}

// Type checking for the declaration
func (d *Declaration) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables, funcEntry *st.FuncEntry) []*SemanticAnalysisError {
	return errors
}

// Get the type of the declaration
func (d *Declaration) GetType() types.Type {
	return d.ty
}

// Get the variable name of the declaration
func (d *Declaration) GetVariable() string {
	return d.variable
}

// Get the token of the declaration
func (d *Declaration) GetToken() *token.Token {
	return d.Token
}

// Control flow for the declaration
func (d *Declaration) GetControlFlow(errors []*SemanticAnalysisError, funcEntry *st.FuncEntry) ([]*SemanticAnalysisError, bool) {
	// Declarations do not have control flow
	return errors, false
}

// Translate the declaration to LLVM IR
func (d *Declaration) ToLLVM(tables *st.SymbolTables) *llvm.GlobalDecl {
	// Declarations are translated in the program and function
	varEntry := tables.Globals.Contains(d.variable)
	return llvm.NewGlobalDecl(d.variable, varEntry.LlvmTy, llvm.GetTypeDefault(d.ty), false)
}
