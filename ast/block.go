package ast

import (
	"bytes"
	st "golite/symboltable"
	"golite/token"
)

// Block node for the AST
type Block struct {
	*token.Token
	statements []Statement
}

// New Block node
func NewBlock(statements []Statement, token *token.Token) *Block {
	return &Block{token, statements}
}

// String representation of the block node
func (b *Block) String() string {
	var out bytes.Buffer

	out.WriteString("{\n")

	for _, stmt := range b.statements {
		out.WriteString(stmt.String())
		out.WriteString("\n")
	}

	out.WriteString("}")

	return out.String()
}

// Build the symbol table for the block
func (b *Block) BuildSymbolTable(tables *st.SymbolTables) {
	// Nothing to do here since we build up the symbol table in the function node
}

// Type check the block
func (b *Block) TypeCheck(errors []string, tables *st.SymbolTables) []string {
	return errors
}
