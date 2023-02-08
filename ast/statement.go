package ast

import st "golite/symboltable"

// Statement is an inferface that all statements implement
// i.e. structs, functions, declarations
type Statement interface {
	String() string
	BuildSymbolTable(tables *st.SymbolTables)
	TypeCheck([]string, *st.SymbolTables) []string
}
