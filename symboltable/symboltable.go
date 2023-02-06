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
	return fmt.Sprintf("[%s (Type: %s) (Scope: %s)]", entry.Name, entry.Ty, scope)
}

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

// Struct entry in the symbol table
type StructEntry struct {
	Name   string
	Fields []*VarEntry
}

// String representation of a struct entry in the symbol table
func (entry *StructEntry) String() string {
	fields := ""
	for _, f := range entry.Fields {
		field := *f
		fields += fmt.Sprintf("%s, ", field.String())
	}
	fields = fields[:len(fields)-2]
	return fmt.Sprintf("%s {Fields: %s}", entry.Name, fields)
}

// Symbol tables for the program (stack)
type SymbolTables struct {
	Globals *SymbolTable[*VarEntry]
	Funcs   *SymbolTable[*FuncEntry]
	Structs *SymbolTable[*StructEntry]
}

// New intance of a stack of symbol tables
func NewSymbolTables() *SymbolTables {
	return &SymbolTables{
		Globals: NewSymbolTable[*VarEntry](nil),
		Funcs:   NewSymbolTable[*FuncEntry](nil),
		Structs: NewSymbolTable[*StructEntry](nil),
	}
}

// Symbol table for a scope
type SymbolTable[T fmt.Stringer] struct {
	parent *SymbolTable[T]
	table  map[string]T
}

// New instance of a symbol table
func NewSymbolTable[T fmt.Stringer](parent *SymbolTable[T]) *SymbolTable[T] {
	return &SymbolTable[T]{
		parent: parent,
		table:  make(map[string]T),
	}
}

// Add an entry to the symbol table
func (st *SymbolTable[T]) Insert(id string, entry T) {
	// Sanity check
	if _, ok := st.table[id]; ok {
		panic(fmt.Sprintf("Symbol %s already exists.", id))
	}
	st.table[id] = entry
}

// Lookup an entry in the symbol table
func (st *SymbolTable[T]) Contains(id string) T {
	if entry, ok := st.table[id]; ok {
		// If the entry is found within the current scope, return it
		return entry
	} else if st.parent != nil {
		// If the entry is not found within the current scope, look in the parent scope (if it exists)
		return st.parent.Contains(id)
	}
	// If the entry is not found within the current scope or the parent scope, return an empty entry
	var emptyEntry T
	return emptyEntry
}
