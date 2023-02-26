package ast

import (
	"golite/llvm"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

// The nil literal
type NilLiteral struct {
	*token.Token            // The token for the integer literal
	Value        string     // The value for the integer literal
	ty           types.Type // The type of the literal
}

// Create a new nil literal
func NewNilLit(token *token.Token) *NilLiteral {
	return &NilLiteral{token, "nil", types.StringToType("nil")}
}

// String representation of the nil literal
func (nl *NilLiteral) String() string {
	return nl.Value
}

// Get the inferred type of the literal
func (nl *NilLiteral) GetType() types.Type {
	return nl.ty
}

// TypeCheck performs type checking for the nil literal
func (nl *NilLiteral) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables, funcEntry *st.FuncEntry) []*SemanticAnalysisError {
	return errors
}

// Translate the nil node into LLVM IR
func (nl *NilLiteral) ToLLVMCFG(tables *st.SymbolTables, blocks []*llvm.BasicBlock, funcEntry *st.FuncEntry) []*llvm.BasicBlock {
	// Update the last block
	block := blocks[len(blocks)-1]

	// Add the integer literal to the last block
	storeInt := llvm.NewStore("null", llvm.GetNextRegister(), "i64")
	// Update the label of the instruction
	storeInt.SetLabel(block.GetLabel())
	block.AddInstruction(storeInt)
	return blocks
}
