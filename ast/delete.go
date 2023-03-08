package ast

import (
	"bytes"
	"golite/llvm"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
	"strings"
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

	// It may be a struct field so we need to split on '.' and check
	exprStr := d.expr.String()
	split := strings.Split(exprStr, ".")
	if len(split) > 1 {
		exprStr = split[0]
	}

	// Check if expression has been declared
	if funcEntry.Variables.Contains(exprStr) == nil {
		// Check the parameters
		found := false
		for _, param := range funcEntry.Parameters {
			if param.Name == exprStr {
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
func (d *Delete) ToLLVMCFG(tables *st.SymbolTables, blocks []*llvm.BasicBlock, funcEntry *st.FuncEntry, constDecls []*llvm.ConstantDecl) ([]*llvm.BasicBlock, []*llvm.ConstantDecl) {
	// Get the LLVM type of the expression
	ty := llvm.TypeToLLVM(d.expr.GetType())
	// Make it a pointer (we know that this will be a struct type)
	ty += "*"

	// Evaluate the expression
	blocks, constDecls, mostRecentOperand := d.expr.ToLLVMCFG(tables, blocks, funcEntry, constDecls)

	// Load the pointer into a register
	loadInst := llvm.NewLoad(mostRecentOperand, ty)
	loadInst.SetLabel(blocks[len(blocks)-1].GetLabel())
	blocks[len(blocks)-1].AddInstruction(loadInst)
	mostRecentOperand = llvm.GetPreviousRegister()

	// Perform a bitcast
	bitcastInst := llvm.NewBitCast(mostRecentOperand, ty, "i8*")
	bitcastInst.SetLabel(blocks[len(blocks)-1].GetLabel())
	blocks[len(blocks)-1].AddInstruction(bitcastInst)

	// Use a free runtime call
	freeInst := llvm.NewFree()
	freeInst.SetLabel(blocks[len(blocks)-1].GetLabel())
	blocks[len(blocks)-1].AddInstruction(freeInst)
	return blocks, constDecls
}
