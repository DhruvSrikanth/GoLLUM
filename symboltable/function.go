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
	for _, param := range entry.Parameters {
		parameters += fmt.Sprintf("%s, ", param.String())
	}
	parameters = parameters[:len(parameters)-2]

	vars := *entry.Variables
	variables := ""
	for k, v := range vars.table {
		variable := *v
		variables += fmt.Sprintf("%s: %s, ", k, variable.String())
	}
	variables = variables[:len(variables)-2]

	return fmt.Sprintf("%s {Return Type: %s} {Parameters: %s} {Variables: %v}", entry.Name, entry.RetTy, parameters, variables)
}
