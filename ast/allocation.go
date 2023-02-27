package ast

import (
	"bytes"
	"golite/llvm"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

// Allocate node in the AST
type Allocate struct {
	*token.Token
	structType string
	ty         types.Type
}

// New Allocate node
func NewAllocate(structType string, token *token.Token) *Allocate {
	ty := types.StringToType("*" + structType)
	return &Allocate{token, structType, ty}
}

// String representation of the allocate node
func (d *Allocate) String() string {
	var out bytes.Buffer

	out.WriteString("new ")
	out.WriteString(d.structType)

	return out.String()
}

// Build the symbol table for the allocate node
func (d *Allocate) BuildSymbolTable(tables *st.SymbolTables, errors []*SemanticAnalysisError) []*SemanticAnalysisError {
	// Nothing to do here since allocate is on a variable thats already been declared and added to the symbol table
	return errors
}

// Type check the allocate node
func (d *Allocate) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables, funcEntry *st.FuncEntry) []*SemanticAnalysisError {
	entry := tables.Structs.Contains(d.structType)
	if entry == nil {
		errors = append(errors, NewSemanticAnalysisError("struct "+d.structType+" not declared", "undeclared struct", d.Token))
	}
	return errors
}

// Get the type of the allocate node
func (d *Allocate) GetType() types.Type {
	// Search for the struct type in the symbol table
	return d.ty
}

// Translate the allocate node into LLVM IR
func (d *Allocate) ToLLVMCFG(tables *st.SymbolTables, blocks []*llvm.BasicBlock, funcEntry *st.FuncEntry) []*llvm.BasicBlock {
	// Get the struct entry
	entry := tables.Structs.Contains(d.structType)
	// Count the number of fields and multiply by 8 to get the size of the struct
	i64Size := 8
	size := len(entry.Fields) * i64Size
	// Allocate the struct
	mallocInst := llvm.NewMalloc(size)
	mallocInst.SetLabel(blocks[len(blocks)-1].GetLabel())
	// Add the malloc instruction to the current block
	blocks[len(blocks)-1].AddInstruction(mallocInst)

	// Get the struct type
	llvmStructType := llvm.TypeToLLVM(d.ty)
	// Cast the malloc result to the struct type using bitcast
	// source register passed in will be the previous register used in the block for the load instruction
	bitcastInst := llvm.NewBitCast(llvm.GetPreviousRegister(), "i8*", llvmStructType)
	bitcastInst.SetLabel(blocks[len(blocks)-1].GetLabel())
	// Add the bitcast instruction to the current block
	blocks[len(blocks)-1].AddInstruction(bitcastInst)

	return blocks
}
