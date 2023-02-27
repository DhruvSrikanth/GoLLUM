package ast

import (
	"fmt"
	"golite/llvm"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
	"strconv"
)

// Integer literal
type IntLiteral struct {
	*token.Token            // The token for the integer literal
	Value        int64      // The value for the integer literal
	ty           types.Type // The type of the literal
}

// Creates a new instance of an integer literal
func NewIntLit(value int64, token *token.Token) *IntLiteral {
	return &IntLiteral{token, value, types.StringToType("int")}
}

// String representation of the integer literal
func (il *IntLiteral) String() string {
	return fmt.Sprintf("%d", il.Value)
}

// Get the inferred type of the string literal
func (il *IntLiteral) GetType() types.Type {
	return il.ty
}

// TypeCheck performs type checking for the integer literal
func (il *IntLiteral) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables, funcEntry *st.FuncEntry) []*SemanticAnalysisError {
	return errors
}

// Translate the allocate node into LLVM IR
func (il *IntLiteral) ToLLVMCFG(tables *st.SymbolTables, blocks []*llvm.BasicBlock, funcEntry *st.FuncEntry) []*llvm.BasicBlock {
	// Add the integer literal to the last block
	storeInt := llvm.NewStore(strconv.Itoa(int(il.Value)), llvm.GetNextRegister(), "i64")
	// Update the label of the instruction
	storeInt.SetLabel(blocks[len(blocks)-1].GetLabel())
	blocks[len(blocks)-1].AddInstruction(storeInt)
	return blocks
}
