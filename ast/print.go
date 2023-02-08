package ast

import (
	"bytes"
	st "golite/symboltable"
	"golite/token"
)

// Print node for the print production rule in the AST
type Print struct {
	*token.Token
	formatString string
	expressions  []Expression // Expressions to be printed
}

// NewPrint node
func NewPrint(formatString string, expressions []Expression, token *token.Token) *Print {
	return &Print{token, formatString, expressions}
}

// String representation of the print node
func (p *Print) String() string {
	var out bytes.Buffer

	out.WriteString("printf(")
	out.WriteString(p.formatString)

	for _, expr := range p.expressions {
		out.WriteString(", ")
		out.WriteString(expr.String())
	}

	out.WriteString(");")

	return out.String()
}

// Build the symbol table for the print
func (p *Print) BuildSymbolTable(tables *st.SymbolTables) {
	// Nothing to do here since a print is not added to the symbol table
}

// Type checking for the print
func (p *Print) TypeCheck(errors []string, tables *st.SymbolTables) []string {
	return errors
}
