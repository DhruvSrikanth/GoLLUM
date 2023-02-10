package symboltable

import (
	"fmt"
)

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
		fmt.Printf("Symbol %s already exists.", id)
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
