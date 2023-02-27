package ast

import (
	"golite/llvm"
	st "golite/symboltable"
	"golite/types"
)

// Expression is the base node for all expression specific nodes (i.e., every expression sub node implements this
// interface)
type Expression interface {
	String() string
	TypeCheck([]*SemanticAnalysisError, *st.SymbolTables, *st.FuncEntry) []*SemanticAnalysisError
	GetType() types.Type
	ToLLVMCFG(*st.SymbolTables, []*llvm.BasicBlock, *st.FuncEntry, []llvm.ConstantDecl) ([]*llvm.BasicBlock, []llvm.ConstantDecl)
}
