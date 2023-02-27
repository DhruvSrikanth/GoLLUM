package ast

import (
	"bytes"
	"golite/llvm"
	st "golite/symboltable"
	"golite/token"
)

// Loop node for the AST
type Loop struct {
	*token.Token
	condition Expression // The condition for the loop
	body      Block      // The block to be executed if the condition is true
}

// NewLoop node
func NewLoop(condition Expression, body Block, token *token.Token) *Loop {
	return &Loop{token, condition, body}
}

// String representation of the loop node
func (l *Loop) String() string {
	var out bytes.Buffer

	out.WriteString("for (")
	out.WriteString(l.condition.String())
	out.WriteString(") ")
	out.WriteString(l.body.String())

	return out.String()
}

// Build the symbol table for the loop node
func (l *Loop) BuildSymbolTable(tables *st.SymbolTables, errors []*SemanticAnalysisError) []*SemanticAnalysisError {
	// Nothing to do here since the body is already in the symbol table
	return errors
}

// Type checking for the loop node
func (l *Loop) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables, funcEntry *st.FuncEntry) []*SemanticAnalysisError {
	// Type check the condition
	errors = l.condition.TypeCheck(errors, tables, funcEntry)

	// Type check the body
	errors = l.body.TypeCheck(errors, tables, funcEntry)

	return errors
}

// Control flow analysis for the loop node
func (l *Loop) GetControlFlow(errors []*SemanticAnalysisError, funcEntry *st.FuncEntry) ([]*SemanticAnalysisError, bool) {
	// Perform control flow analysis on the block
	var flow bool
	errors, flow = l.body.GetControlFlow(errors, funcEntry)
	return errors, flow
}

// Translate the loop node to LLVM IR
func (l *Loop) ToLLVMCFG(tables *st.SymbolTables, blocks []*llvm.BasicBlock, funcEntry *st.FuncEntry, constDecls []*llvm.ConstantDecl) ([]*llvm.BasicBlock, []*llvm.ConstantDecl) {
	// Add new block for the condition
	condBlock := llvm.NewBasicBlock(llvm.GetNextLabel())
	// Add the condition instructions to the block
	blocks = append(blocks, condBlock)
	// Get the condition expression to translate to LLVM IR
	// blocks := l.condition.ToLLVM(tables, blocks, funcEntry)

	// Add new block for the body block
	bodyBlock := llvm.NewBasicBlock(llvm.GetNextLabel())
	// Append the then block to the list of blocks since the last block will be considered for the block to add instructions to
	blocks = append(blocks, bodyBlock)
	blocks, constDecls = l.body.ToLLVMCFG(tables, blocks, funcEntry, constDecls)

	// Add new block for canonical form
	block := llvm.NewBasicBlock(llvm.GetNextLabel())
	// Add the block to the list of blocks
	blocks = append(blocks, block)

	return blocks, constDecls
}
