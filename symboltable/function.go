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
		if i < len(entry.Parameters)-1 {
			parameters += fmt.Sprintf("%s, ", param.String())
		} else {
			parameters += param.String()
		}
	}

	vars := *entry.Variables
	variables := ""
	i := 0
	for k, v := range vars.table {
		variable := *v
		if i < len(vars.table)-1 {
			variables += fmt.Sprintf("%s: %s, ", k, variable.String())
		} else {
			variables += fmt.Sprintf("%s: %s", k, variable.String())
		}
	}

	return fmt.Sprintf("%s {Return Type: %s} {Parameters: %s} {Variables: %v}", entry.Name, entry.RetTy, parameters, variables)
}
