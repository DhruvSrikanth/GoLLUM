package ast

import (
	"bytes"
	"golite/llvm"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
	"strings"
)

// Print node for the print production rule in the AST
type Print struct {
	*token.Token
	formatString string
	expressions  []Expression // Expressions to be printed
}

// NewPrint node
func NewPrint(formatString string, expressions []Expression, token *token.Token) *Print {
	return &Print{token, formatString, expressions}
}

// String representation of the print node
func (p *Print) String() string {
	var out bytes.Buffer

	out.WriteString("printf(")
	out.WriteString(p.formatString)

	for _, expr := range p.expressions {
		out.WriteString(", ")
		out.WriteString(expr.String())
	}

	out.WriteString(");")

	return out.String()
}

// Build the symbol table for the print
func (p *Print) BuildSymbolTable(tables *st.SymbolTables, errors []*SemanticAnalysisError) []*SemanticAnalysisError {
	// Nothing to do here since a print is not added to the symbol table
	return errors
}

// Type checking for the print
func (p *Print) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables, funcEntry *st.FuncEntry) []*SemanticAnalysisError {
	// Type check all of the expressions in the print
	// this ensures that GetType() on any expression will be valid
	for _, expr := range p.expressions {
		errors = expr.TypeCheck(errors, tables, funcEntry)
	}

	// Check if the format string contains the same number of expressions as the number of %d
	if len(p.expressions) != strings.Count(p.formatString, "%d") {
		errors = append(errors, NewSemanticAnalysisError("Number of expressions does not match the number of %d in the format string", "placeholder mismatch", p.Token))
	}

	// Check that the expressions are all of type int since that is the only placeholder allowed
	for _, expr := range p.expressions {
		if expr.GetType() != types.StringToType("int") {
			errors = append(errors, NewSemanticAnalysisError("Expression is not of type int", "invalid type", p.Token))
		}
	}

	return errors
}

// Control flow analysis for the print node
func (p *Print) GetControlFlow(errors []*SemanticAnalysisError, funcEntry *st.FuncEntry) ([]*SemanticAnalysisError, bool) {
	// Nothing to do here since the print node does not have any control flow
	return errors, false
}

// Translate the print node to LLVM IR
func (p *Print) ToLLVMCFG(tables *st.SymbolTables, blocks []*llvm.BasicBlock, funcEntry *st.FuncEntry, constDecls []*llvm.ConstantDecl) ([]*llvm.BasicBlock, []*llvm.ConstantDecl) {
	// Stay in the same block
	size := len(p.formatString)
	// Add 1 for the null terminator
	size++
	// Subtract 2 for the " at the start and end
	size -= 2
	// Format the format string
	// Replace all %d with %ld
	formatString := strings.Replace(p.formatString, "%d", "%ld", -1)
	// Replace all \n with \0A
	formatString = strings.Replace(formatString, "\n", "\\0A", -1)
	// No need to to add a null terminator since that is added as part of the constant decl

	// Create the constant decl
	// constDecl := llvm.NewConstantDecl("format_string", llvm.NewArrayType(size, llvm.NewIntType(8)), llvm.NewStringConstant(formatString))
	// constDecls = append(constDecls, constDecl)

	return blocks, constDecls
}
