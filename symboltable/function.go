package symboltable

import (
	"fmt"
	"golite/types"
)

// Function entry in the symbol table
type FuncEntry struct {
	Name       string
	RetTy      types.Type
	Parameters []*VarEntry
	Variables  *SymbolTable[*VarEntry]
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
	out += fmt.Sprintf("Return Type -  %s\n", entry.RetTy.String())
	if len(entry.Parameters) > 0 {
		out += fmt.Sprintf("Parameters -  %s\n", parameters)
	}
	if len(entry.Variables.table) > 0 {
		out += fmt.Sprintf("Variables -  %s\n", variables)
	}

	return out
}
