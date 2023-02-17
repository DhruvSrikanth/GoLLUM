package symboltable

import (
	"fmt"
	"golite/types"
)

// Type Aliasing for a variable entry in the symbol table
type VarScope int

// Scope of a variable
const (
	GLOBAL VarScope = iota
	LOCAL
)

// Variable entry in the symbol table
type VarEntry struct {
	Name  string
	Ty    types.Type
	Scope VarScope
}

// String representation of an variable entry in the symbol table
func (entry *VarEntry) String() string {
	scope := "Global"
	if entry.Scope == LOCAL {
		scope = "Local"
	}
	return fmt.Sprintf("%s (Type: %s) (Scope: %s)", entry.Name, entry.Ty, scope)
}

// Get the type of the variable entry
func (entry *VarEntry) GetType() types.Type {
	return entry.Ty
}
