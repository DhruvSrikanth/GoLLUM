package ast

import (
	"fmt"
	"golite/llvm"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

// Boolean literal
type BoolLiteral struct {
	*token.Token            // The token for the integer literal
	Value        bool       // The value for the integer literal
	ty           types.Type // The type of the literal
}

// Creates a new instance of a booleal literal
func NewBoolLit(value bool, token *token.Token) *BoolLiteral {
	return &BoolLiteral{token, value, types.StringToType("bool")}
}

// String representation of the bool literal
func (bl *BoolLiteral) String() string {
	return fmt.Sprintf("%t", bl.Value)
}

// Get the inferred type of the literal
func (bl *BoolLiteral) GetType() types.Type {
	return bl.ty
}

// TypeCheck performs type checking for the bool literal
func (bl *BoolLiteral) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables, funcEntry *st.FuncEntry) []*SemanticAnalysisError {
	return errors
}

// Translate the bool literal into LLVM IR
func (bl *BoolLiteral) ToLLVMCFG(tables *st.SymbolTables, blocks []*llvm.BasicBlock, funcEntry *st.FuncEntry) []*llvm.BasicBlock {
	// Add the integer literal to the last block
	var boolVal string
	if bl.Value {
		boolVal = "1"
	} else {
		boolVal = "0"
	}
	storeInt := llvm.NewStore(boolVal, llvm.GetNextRegister(), "i64")
	// Update the label of the instruction
	storeInt.SetLabel(blocks[len(blocks)-1].GetLabel())
	blocks[len(blocks)-1].AddInstruction(storeInt)
	return blocks
}
