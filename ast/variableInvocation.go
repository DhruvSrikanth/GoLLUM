package ast

import (
	"bytes"
	"fmt"
	"golite/llvm"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
	"strings"
)

// VariableInvocation node for the AST
type VariableInvocation struct {
	*token.Token
	identifier string       // The identifier of the variable being invoked
	arguments  []Expression // The arguments to the variable being invoked (function call case)
	ty         types.Type   // The type of the variable being invoked (or return type of the function being called)
}

// NewVariableInvocation node
func NewVariableInvocation(identifier string, arguments []Expression, token *token.Token) *VariableInvocation {
	return &VariableInvocation{token, identifier, arguments, nil}
}

// String representation of the variable invocation node
func (v *VariableInvocation) String() string {
	var out bytes.Buffer

	out.WriteString(v.identifier)
	if len(v.arguments) > 0 {
		out.WriteString("(")
		for x, arg := range v.arguments {
			out.WriteString(arg.String())
			if x < len(v.arguments)-1 {
				out.WriteString(", ")
			}
		}
		out.WriteString(")")
	}

	return out.String()
}

// Build the symbol table for the VariableInvocation node
func (v *VariableInvocation) BuildSymbolTable(tables *st.SymbolTables, errors []*SemanticAnalysisError) []*SemanticAnalysisError {
	// Nothing to do here since the function is already in the symbol table and the VariableInvocation is not a declaration
	return errors
}

// Type checking for the VariableInvocation node
func (v *VariableInvocation) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables, funcEntry *st.FuncEntry) []*SemanticAnalysisError {
	// Check if its a variable or function call
	if len(v.arguments) != 0 {
		// type check the arguments
		for _, arg := range v.arguments {
			errors = arg.TypeCheck(errors, tables, funcEntry)
		}

		// Check if the function exists
		entry := tables.Funcs.Contains(v.identifier)
		if entry == nil {
			errors = append(errors, NewSemanticAnalysisError(fmt.Sprintf("function %s does not exist", v.identifier), "undeclared function", v.Token))
		} else {
			// Check if the number of arguments is correct
			if len(v.arguments) != len(entry.Parameters) {
				errors = append(errors, NewSemanticAnalysisError(fmt.Sprintf("function %s expects %d arguments, got %d", v.identifier, len(entry.Parameters), len(v.arguments)), "incorrect number of arguments", v.Token))
			} else {
				// Check if the arguments are of the correct type
				for x, arg := range v.arguments {
					// fmt.Println(arg.GetType(), arg)
					if arg.GetType() != entry.Parameters[x].GetType() {
						errors = append(errors, NewSemanticAnalysisError(fmt.Sprintf("function %s expects %s as argument %d, got %s", v.identifier, entry.Parameters[x].GetType(), x+1, arg.GetType()), "mismatched type", v.Token))
					}
				}
			}

			// Set the type of the VariableInvocation
			v.ty = entry.RetTy
		}

	} else {
		// Variable
		entry := funcEntry.Variables.Contains(v.identifier)
		// Search through the parameters
		if entry == nil {
			for _, param := range funcEntry.Parameters {
				if param.Name == v.identifier {
					entry = param
					break
				}
			}
		}

		// Check if the variable exists in the symbol table
		if entry == nil {
			errors = append(errors, NewSemanticAnalysisError("variable "+v.identifier+" not declared.", "undeclared variable", v.Token))
		}

		// Set the type of the VariableInvocation
		// This will throw an immediate error if the variable is not declared
		v.ty = types.StringToType("nil")
	}

	return errors
}

// Get the type of the VariableInvocation node
func (v *VariableInvocation) GetType() types.Type {
	return v.ty
}

// Translate the allocate node into LLVM IR
func (v *VariableInvocation) ToLLVMCFG(tables *st.SymbolTables, blocks []*llvm.BasicBlock, funcEntry *st.FuncEntry, constDecls []*llvm.ConstantDecl) ([]*llvm.BasicBlock, []*llvm.ConstantDecl, string) {
	// Check if its a variable or function call
	var mostRecentOperand string
	if len(v.arguments) != 0 {
		// Function
		// Get the function entry
		entry := tables.Funcs.Contains(v.identifier)
		// Need to load all of the arguments into registers
		// Record the last instruction source register in the block
		argRegs := make([]string, 0)
		for _, param := range v.arguments {
			// Load the argument into a register by calling the ToLLVMCFG function
			blocks, constDecls, mostRecentOperand = param.ToLLVMCFG(tables, blocks, funcEntry, constDecls)
			// If the most recent operand is a constant, we need to store it into a register
			if !strings.HasPrefix(mostRecentOperand, "%r") {
				storeInst := llvm.NewStore(mostRecentOperand, llvm.GetNextRegister(), "i64")
				storeInst.SetLabel(blocks[len(blocks)-1].GetLabel())
				blocks[len(blocks)-1].AddInstruction(storeInst)
				mostRecentOperand = llvm.GetPreviousRegister()
			} else if strings.Contains(blocks[len(blocks)-1].GetLastInstruction().String(), "getelementptr") {
				// If the most recent operand is a pointer, we need to load it into a register
				loadInst := llvm.NewLoad(mostRecentOperand, "i64")
				loadInst.SetLabel(blocks[len(blocks)-1].GetLabel())
				blocks[len(blocks)-1].AddInstruction(loadInst)
				mostRecentOperand = llvm.GetPreviousRegister()
			}

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
		mostRecentOperand = llvm.GetPreviousRegister()
	} else {
		// Variable
		// Get the variable entry
		entry := funcEntry.Variables.Contains(v.identifier)
		if entry == nil {
			for _, param := range funcEntry.Parameters {
				if param.Name == v.identifier {
					entry = param
					break
				}
			}

			// Load the parameter into a register
			// Make the load instruction
			ty := entry.LlvmTy
			if strings.Contains(ty, "struct.") {
				ty += "*"
			}
			loadInst := llvm.NewLoad("%P_"+entry.Name, ty)
			loadInst.SetLabel(blocks[len(blocks)-1].GetLabel())
			blocks[len(blocks)-1].AddInstruction(loadInst)
			mostRecentOperand = llvm.GetPreviousRegister()

		} else {
			// Load the variable into a register
			var varName string
			if entry.Scope == st.GLOBAL {
				varName = "@" + entry.Name
			} else {
				varName = "%" + entry.Name
			}

			ty := entry.LlvmTy
			if strings.Contains(ty, "struct.") {
				ty += "*"
			}

			loadInst := llvm.NewLoad(varName, ty)
			loadInst.SetLabel(blocks[len(blocks)-1].GetLabel())
			blocks[len(blocks)-1].AddInstruction(loadInst)
			mostRecentOperand = llvm.GetPreviousRegister()
		}
	}
	return blocks, constDecls, mostRecentOperand
}
