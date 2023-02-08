package ast

import (
	st "golite/symboltable"
	"golite/types"
)

// Expression is the base node for all expression specific nodes (i.e., every expression sub node implements this
// interface)
type Expression interface {
	String() string
	TypeCheck([]string, *st.SymbolTables) []string
	GetType() types.Type
}
