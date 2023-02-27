package ast

import (
	"bytes"
	"golite/llvm"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
	"strings"
)

// Return production rule node in the AST
type Return struct {
	*token.Token
	expression *Expression // The expression to be returned
	ty         types.Type  // The type of the expression
}

// NewReturn node
func NewReturn(expression *Expression, token *token.Token) *Return {
	return &Return{token, expression, nil}
}

// String representation of the return node
func (r *Return) String() string {
	var out bytes.Buffer

	out.WriteString("return ")
	if *r.expression != nil {
		out.WriteString((*r.expression).String())
	}
	out.WriteString(";")

	return out.String()
}

// Build the symbol table for the return node
func (r *Return) BuildSymbolTable(tables *st.SymbolTables, errors []*SemanticAnalysisError) []*SemanticAnalysisError {
	// Nothing to do here since the expression variables are already in the symbol table
	return errors
}

// Type checking for the return node
func (r *Return) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables, funcEntry *st.FuncEntry) []*SemanticAnalysisError {
	// If the expression does not exist, then the return type is void
	if *r.expression == nil {
		r.ty = types.StringToType("void")
	} else {
		// Type check the expression
		// Ensures that GetType() has a valid type
		errors = (*r.expression).TypeCheck(errors, tables, funcEntry)

		// Get the type of the expression
		r.ty = (*r.expression).GetType()
	}

	// Check if the return type is the same as the function return type
	if r.ty != funcEntry.RetTy {
		errors = append(errors, NewSemanticAnalysisError("return type does not match function signature return type", "mismatched types", r.Token))
	}

	return errors
}

// Get the type of the return node
func (r *Return) GetType() types.Type {
	return r.ty
}

// Control flow analysis for the return node
func (r *Return) GetControlFlow(errors []*SemanticAnalysisError, funcEntry *st.FuncEntry) ([]*SemanticAnalysisError, bool) {
	// Return true to indicate that the return node has control flow
	return errors, true
}

// Translate the return node to LLVM IR
func (r *Return) ToLLVMCFG(tables *st.SymbolTables, blocks []*llvm.BasicBlock, funcEntry *st.FuncEntry, constDecls []llvm.ConstantDecl) ([]*llvm.BasicBlock, []llvm.ConstantDecl) {
	// If the expression exists, then translate it
	if *r.expression != nil {
		// Translate the expression
		blocks, constDecls = (*r.expression).ToLLVMCFG(tables, blocks, funcEntry, constDecls)
		// If this exists, then store the result from the last used register into a new register
		lastUsedReg := llvm.GetPreviousRegister()
		exprLLVMType := llvm.TypeToLLVM(r.ty)
		if strings.Contains(exprLLVMType, "struct.") {
			exprLLVMType += "*"
		}

		storeInst := llvm.NewStore(lastUsedReg, "%_retval", exprLLVMType)
		storeInst.SetLabel(blocks[len(blocks)-1].GetLabel())
		blocks[len(blocks)-1].AddInstruction(storeInst)

		// Load the return value into the a register
		loadInst := llvm.NewLoad("%_retval", exprLLVMType)
		loadInst.SetLabel(blocks[len(blocks)-1].GetLabel())
		blocks[len(blocks)-1].AddInstruction(loadInst)

		// Create the return instruction
		retInst := llvm.NewReturn("%_retval", exprLLVMType)
		retInst.SetLabel(blocks[len(blocks)-1].GetLabel())
		blocks[len(blocks)-1].AddInstruction(retInst)
	} else {
		// Expression does not exist, so we just return a 0 since we treat this as a void function
		// the void function is represented by i64 in LLVM IR
		// First store a 0 into the return value register
		storeInst := llvm.NewStore("0", "%_retval", "i64")
		storeInst.SetLabel(blocks[len(blocks)-1].GetLabel())
		blocks[len(blocks)-1].AddInstruction(storeInst)

		// Load the return value into the a register
		loadInst := llvm.NewLoad("%_retval", "i64")
		loadInst.SetLabel(blocks[len(blocks)-1].GetLabel())
		blocks[len(blocks)-1].AddInstruction(loadInst)

		// Create the return instruction
		retInst := llvm.NewReturn("%_retval", "i64")
		retInst.SetLabel(blocks[len(blocks)-1].GetLabel())
		blocks[len(blocks)-1].AddInstruction(retInst)
	}
	return blocks, constDecls
}
