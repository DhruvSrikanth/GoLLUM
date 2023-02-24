package ast

import (
	"golite/llvm"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

// DeclarationStatement is an interface for all functions
type DeclarationStatement interface {
	Statement
	GetType() types.Type
	GetVariable() string
	GetToken() *token.Token
	ToLLVM(*st.SymbolTables) *llvm.GlobalDecl
}
