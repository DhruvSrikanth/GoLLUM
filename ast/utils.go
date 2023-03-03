package ast

import (
	st "golite/symboltable"
)

// Is the variable a named variable
func isNamed(funcEntry *st.FuncEntry, id string) bool {
	if id == "" {
		return false
	}

	if funcEntry.Variables.Contains(id) != nil {
		return true
	}

	for _, param := range funcEntry.Parameters {
		if param.Name == id || "P_"+param.Name == id {
			return true
		}
	}

	return false
}
