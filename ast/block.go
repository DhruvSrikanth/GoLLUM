package ast

import (
	"bytes"
	"golite/llvm"
	st "golite/symboltable"
	"golite/token"
)

// Block node for the AST
type Block struct {
	*token.Token
	statements []Statement
}

// New Block node
func NewBlock(statements []Statement, token *token.Token) *Block {
	return &Block{token, statements}
}

// String representation of the block node
func (b *Block) String() string {
	var out bytes.Buffer

	out.WriteString("{\n")

	for _, stmt := range b.statements {
		out.WriteString(stmt.String())
		out.WriteString("\n")
	}

	out.WriteString("}")

	return out.String()
}

// Build the symbol table for the block
func (b *Block) BuildSymbolTable(tables *st.SymbolTables, errors []*SemanticAnalysisError) []*SemanticAnalysisError {
	// Nothing to do here since we build up the symbol table in the function node
	return errors
}

// Type check the block
func (b *Block) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables, funcEntry *st.FuncEntry) []*SemanticAnalysisError {
	for _, stmt := range b.statements {
		errors = stmt.TypeCheck(errors, tables, funcEntry)
	}
	return errors
}

// Control flow analysis for the block node
func (b *Block) GetControlFlow(errors []*SemanticAnalysisError, funcEntry *st.FuncEntry) ([]*SemanticAnalysisError, bool) {
	// Nothing to do here since the block node does not have any control flow
	flow := false
	cflow := false
	for _, stmt := range b.statements {
		errors, cflow = stmt.GetControlFlow(errors, funcEntry)
		flow = flow || cflow
	}
	return errors, flow
}

// Translate the block node to LLVM IR
func (b *Block) ToLLVMCFG(tables *st.SymbolTables, blocks []*llvm.BasicBlock, funcEntry *st.FuncEntry) []*llvm.BasicBlock {
	// Translate each statement in the block
	for _, stmt := range b.statements {
		blocks = stmt.ToLLVMCFG(tables, blocks, funcEntry)
	}
	return blocks
}
