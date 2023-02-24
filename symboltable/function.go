package symboltable

import (
	"fmt"
	"golite/llvm"
	"golite/types"
)

// Function entry in the symbol table
type FuncEntry struct {
	Name       string
	RetTy      types.Type
	Parameters []*VarEntry
	Variables  *SymbolTable[*VarEntry]
	LlvmRetTy  string
}

// String representation of a function entry in the symbol table
func (entry *FuncEntry) String() string {
	parameters := ""
	for i, param := range entry.Parameters {
		if i < len(entry.Parameters)-1 && len(entry.Parameters) > 1 {
			parameters += fmt.Sprintf("%s, ", param.String())
		} else {
			parameters += param.String()
		}
	}

	vars := *entry.Variables
	variables := ""
	i := 0
	for _, v := range vars.table {
		variable := *v
		if i < len(vars.table)-1 && len(vars.table) > 1 {
			variables += fmt.Sprintf("%s, ", variable.String())
		} else {
			variables += variable.String()
		}
	}
	out := entry.Name + "\n"
	out += fmt.Sprintf("Return Type - %s [%s]\n", entry.RetTy.String(), entry.LlvmRetTy)
	if len(entry.Parameters) > 0 {
		out += fmt.Sprintf("Parameters -  %s\n", parameters)
	}
	if len(entry.Variables.table) > 0 {
		out += fmt.Sprintf("Variables -  %s\n", variables)
	}

	return out
}

// Get the type of the function
func (entry *FuncEntry) GetType() types.Type {
	return entry.RetTy
}

// Translate the function entry to its LLVM representation
func (entry *FuncEntry) LLVMTypeConversion() {
	// Set the return type
	entry.LlvmRetTy = llvm.TypeToLLVM(entry.RetTy)

	// Set the parameter types
	for _, param := range entry.Parameters {
		param.LLVMTypeConversion()
	}

	// Set the variable types
	for _, v := range entry.Variables.table {
		v.LLVMTypeConversion()
	}
}
