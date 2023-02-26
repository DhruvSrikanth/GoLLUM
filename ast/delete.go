package ast

import (
	"bytes"
	"golite/llvm"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

// Delete node in the AST
type Delete struct {
	*token.Token
	expr Expression // Expression to be deleted
}

// New Delete node
func NewDelete(expr Expression, token *token.Token) *Delete {
	return &Delete{token, expr}
}

// String representation of the delete node
func (d *Delete) String() string {
	var out bytes.Buffer

	out.WriteString("delete ")
	out.WriteString(d.expr.String())
	out.WriteString(";")

	return out.String()
}

// Build the symbol table for the delete node
func (d *Delete) BuildSymbolTable(tables *st.SymbolTables, errors []*SemanticAnalysisError) []*SemanticAnalysisError {
	// Nothing to do here since a delete is not added to the symbol table
	return errors
}

// Type check the delete node
func (d *Delete) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables, funcEntry *st.FuncEntry) []*SemanticAnalysisError {
	// Type check the expression
	errors = d.expr.TypeCheck(errors, tables, funcEntry)

	// Check if expression has been declared
	if funcEntry.Variables.Contains(d.expr.String()) == nil {
		// Check the parameters
		found := false
		for _, param := range funcEntry.Parameters {
			if param.Name == d.expr.String() {
				found = true
			}
		}
		if !found {
			errors = append(errors, NewSemanticAnalysisError(d.expr.String()+" has not been declared.", "undeclared identifier", d.Token))
		}
	}

	// Check if the expression is a pointer to a struct type
	exprType := d.expr.GetType()
	if types.TypeToKind(exprType) != types.STRUCT {
		errors = append(errors, NewSemanticAnalysisError("delete can only be called on pointers to non-primitive types.", "invalid memory management", d.Token))
	} else if tables.Structs.Contains(exprType.String()[1:]) == nil {
		// Check if struct declared in symbol table
		errors = append(errors, NewSemanticAnalysisError("struct "+exprType.String()[1:]+" does not exist.", "invalid type", d.Token))
	}
	return errors
}

// Control flow analysis for the delete node
func (d *Delete) GetControlFlow(errors []*SemanticAnalysisError, funcEntry *st.FuncEntry) ([]*SemanticAnalysisError, bool) {
	// Nothing to do here since the delete node does not have any control flow
	return errors, false
}

// Translate the delete node to LLVM IR
func (d *Delete) ToLLVMCFG(tables *st.SymbolTables, blocks []*llvm.BasicBlock, funcEntry *st.FuncEntry) []*llvm.BasicBlock {
	// Stay in the same block
	block := blocks[len(blocks)-1]

	// Add load instruction to the block
	// Get the variable entry
	localVariable := funcEntry.Variables.Contains(d.expr.String())
	if localVariable == nil {
		// Check the parameters
		for _, param := range funcEntry.Parameters {
			if param.Name == d.expr.String() {
				localVariable = param
			}
		}
	}
	loadInst := llvm.NewLoad(localVariable.Name, localVariable.LlvmTy)
	// Update the instruction label
	loadInst.SetLabel(block.GetLabel())
	// Add the instruction to the block
	block.AddInstruction(loadInst)

	// Add the bitcast instruction to the block
	bitcastInst := llvm.NewBitCast(localVariable.LlvmTy, "i8")
	// // Update the instruction label
	bitcastInst.SetLabel(block.GetLabel())
	// // Add the instruction to the block
	block.AddInstruction(bitcastInst)

	// Add the free runtime function call instruction to the block
	// freeInst := llvm.NewCall("free", []llvm.Value{loadInst.GetResult()})
	// // Update the instruction label
	// freeInst.SetLabel(block.GetLabel())
	// // Add the instruction to the block
	// block.AddInstruction(freeInst)

	// Update the same block with the assignment
	blocks[len(blocks)-1] = block
	return blocks
}
