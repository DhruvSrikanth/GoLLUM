package symboltable

import (
	"bytes"
	"fmt"
	"strconv"
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

// String representation of the symbol tables
func (tables *SymbolTables) String() string {
	var out bytes.Buffer
	out.WriteString("Symbol Tables:\n")
	out.WriteString("-------------\n")
	out.WriteString("Structs - \n")
	out.WriteString(tables.Structs.String())
	out.WriteString("Globals - \n")
	out.WriteString(tables.Globals.String())
	out.WriteString("Functions - \n")
	out.WriteString(tables.Funcs.String())
	out.WriteString("-------------\n")
	return out.String()
}

// Translate the symbol tables to its LLVM representation
func (tables *SymbolTables) LLVMTypeConversion() {
	// Convert the structs
	for _, v := range tables.Structs.table {
		v.LLVMTypeConversion()
	}
	// Convert the functions
	for _, v := range tables.Funcs.table {
		v.LLVMTypeConversion()
	}
	// Convert the globals
	for _, v := range tables.Globals.table {
		v.LLVMTypeConversion()
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

// String representation of the symbol table
func (s *SymbolTable[T]) String() string {
	var out bytes.Buffer

	out.WriteString("=============\n")
	i := 0
	for _, v := range s.table {
		i += 1
		out.WriteString(strconv.Itoa(i) + ". " + v.String() + "\n")
	}
	out.WriteString("=============\n")

	return out.String()
}

// Add an entry to the symbol table
func (st *SymbolTable[T]) Insert(id string, entry T) bool {
	redeclaration := true

	// Sanity check
	if _, ok := st.table[id]; ok {
		fmt.Printf("Symbol %s already exists.", id)
		redeclaration = false
	}
	st.table[id] = entry
	return redeclaration
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

// Get the table of the symbol table
func (st *SymbolTable[T]) GetTable() map[string]T {
	return st.table
}
