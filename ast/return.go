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
func (r *Return) ToLLVMCFG(tables *st.SymbolTables, blocks []*llvm.BasicBlock, funcEntry *st.FuncEntry, constDecls []*llvm.ConstantDecl) ([]*llvm.BasicBlock, []*llvm.ConstantDecl) {
	var mostRecentOperand string
	var llvmTy string

	// If the expression exists, then translate it
	if *r.expression != nil {
		// If the expression is not void type, then translate it
		// Translate the expression
		blocks, constDecls, mostRecentOperand = (*r.expression).ToLLVMCFG(tables, blocks, funcEntry, constDecls)

		llvmTy = llvm.TypeToLLVM(r.ty)
		if strings.Contains(llvmTy, "struct.") {
			llvmTy += "*"
		}

		if isNamed(funcEntry, mostRecentOperand[1:]) || (strings.Contains((*r.expression).String(), ".") && !strings.Contains((*r.expression).String(), " ")) {
			// Load the expression
			loadInst := llvm.NewLoad(mostRecentOperand, llvmTy)
			loadInst.SetLabel(blocks[len(blocks)-1].GetLabel())
			blocks[len(blocks)-1].AddInstruction(loadInst)
			mostRecentOperand = llvm.GetPreviousRegister()
		}

	} else {
		// If the expression does not exist, return the void type (i64 in LLVM IR)
		llvmTy = "i64"
		// Expression does not exist, so we just return a 0 since we treat this as a void function (i64 LLVM IR - default value is 0)
		mostRecentOperand = "0"
	}
	// Perform a store instruction
	storeInst := llvm.NewStore(mostRecentOperand, "%"+funcEntry.Name+"_retval", llvmTy)
	storeInst.SetLabel(blocks[len(blocks)-1].GetLabel())
	blocks[len(blocks)-1].AddInstruction(storeInst)

	// Load the return value into the a register
	loadInst := llvm.NewLoad("%"+funcEntry.Name+"_retval", llvmTy)
	loadInst.SetLabel(blocks[len(blocks)-1].GetLabel())
	blocks[len(blocks)-1].AddInstruction(loadInst)
	mostRecentOperand = llvm.GetPreviousRegister()

	// Create the return instruction
	retInst := llvm.NewReturn(mostRecentOperand, llvmTy)
	retInst.SetLabel(blocks[len(blocks)-1].GetLabel())
	blocks[len(blocks)-1].AddInstruction(retInst)

	return blocks, constDecls
}
