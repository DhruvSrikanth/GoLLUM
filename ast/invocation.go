package ast

import (
	"bytes"
	"golite/llvm"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
	"strings"
)

// Invocation node for the AST
type Invocation struct {
	*token.Token
	identifier string       // The identifier of the function being invoked
	arguments  []Expression // The arguments to the function
	ty         types.Type
}

// NewInvocation node
func NewInvocation(identifier string, arguments []Expression, token *token.Token) *Invocation {
	return &Invocation{token, identifier, arguments, nil}
}

// String representation of the invocation node
func (i *Invocation) String() string {
	var out bytes.Buffer

	out.WriteString(i.identifier)
	out.WriteString("(")
	for x, arg := range i.arguments {
		out.WriteString(arg.String())
		if x < len(i.arguments)-1 {
			out.WriteString(", ")
		}
	}
	out.WriteString(")")
	out.WriteString(";")

	return out.String()
}

// Build the symbol table for the invocation node
func (i *Invocation) BuildSymbolTable(tables *st.SymbolTables, errors []*SemanticAnalysisError) []*SemanticAnalysisError {
	// Nothing to do here since the function is already in the symbol table and the invocation is not a declaration
	return errors
}

// Type checking for the invocation node
func (i *Invocation) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables, funcEntry *st.FuncEntry) []*SemanticAnalysisError {
	// Type check the expressions
	// Ensures that GetType() has a valid type
	for _, arg := range i.arguments {
		errors = arg.TypeCheck(errors, tables, funcEntry)
	}

	// Check if the function has been declared
	entry := tables.Funcs.Contains(i.identifier)
	if entry == nil {
		errors = append(errors, NewSemanticAnalysisError("function "+i.identifier+" has not been declared", "Undeclared function", i.Token))
	} else {
		// Check if the number of arguments is correct
		if len(i.arguments) != len(entry.Parameters) {
			errors = append(errors, NewSemanticAnalysisError("function "+i.identifier+" has an incorrect number of arguments passed", "incorrect arguments provided", i.Token))
		} else {
			// Check the types of the arguments
			// Can combine the with the above loop, however, this is more readable and explicit
			for x, arg := range i.arguments {
				if arg.GetType() != entry.Parameters[x].GetType() {
					errors = append(errors, NewSemanticAnalysisError("argument type mismatch in function call", "mismatched type", i.GetToken()))
				}
			}
		}
	}

	// Set the type of the invocation node
	i.ty = funcEntry.RetTy

	return errors
}

// Get the type of the invocation node
func (i *Invocation) GetType() types.Type {
	return i.ty
}

// Get the token of the invocation node
func (i *Invocation) GetToken() *token.Token {
	return i.Token
}

// Control flow analysis for the invocation node
func (i *Invocation) GetControlFlow(errors []*SemanticAnalysisError, funcEntry *st.FuncEntry) ([]*SemanticAnalysisError, bool) {
	// Nothing to do here since the invocation node does not have any control flow
	return errors, false
}

// Translate the invocation node to LLVM IR
func (i *Invocation) ToLLVMCFG(tables *st.SymbolTables, blocks []*llvm.BasicBlock, funcEntry *st.FuncEntry, constDecls []*llvm.ConstantDecl) ([]*llvm.BasicBlock, []*llvm.ConstantDecl) {
	// Stay in the same block
	var mostRecentOperand string
	// Function
	// Get the function entry
	entry := tables.Funcs.Contains(i.identifier)
	// Need to load all of the arguments into registers
	// Record the last instruction source register in the block
	argRegs := make([]string, 0)
	for _, param := range i.arguments {
		// Load the argument into a register by calling the ToLLVMCFG function
		blocks, constDecls, mostRecentOperand = param.ToLLVMCFG(tables, blocks, funcEntry, constDecls)
		// Get the last used register
		argRegs = append(argRegs, mostRecentOperand)
	}

	argTypes := make([]string, 0)
	for _, param := range entry.Parameters {
		paramTy := param.LlvmTy
		if strings.Contains(paramTy, "struct.") {
			paramTy += "*"
		}
		argTypes = append(argTypes, param.LlvmTy)
	}

	retTy := entry.LlvmRetTy
	if strings.Contains(retTy, "struct.") {
		retTy += "*"
	}

	// Make the function call
	// Note that the reassignment of the return to a variable is done in the assignment node as a store instruction
	fCallInst := llvm.NewFunctionCall(entry.Name, retTy, argRegs, argTypes)
	fCallInst.SetLabel(blocks[len(blocks)-1].GetLabel())
	blocks[len(blocks)-1].AddInstruction(fCallInst)
	return blocks, constDecls
}
