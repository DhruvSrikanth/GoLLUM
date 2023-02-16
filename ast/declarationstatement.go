package ast

import "golite/types"

// DeclarationStatement is an interface for all functions
type DeclarationStatement interface {
	Statement
	GetType() types.Type
	GetVariable() string
}
