package ast

import (
	"bytes"
	st "golite/symboltable"
	"golite/token"
)

// Conditional statement node for the AST
type Conditional struct {
	*token.Token
	condition Expression // The condition for the if statement
	thenStmt  Block      // The block to be executed if the condition is true
	elseStmt  Block      // The block to be executed if the condition is false
}

// New Conditional node
func NewConditional(condition Expression, thenStmt Block, elseStmt Block, token *token.Token) *Conditional {
	return &Conditional{token, condition, thenStmt, elseStmt}
}

// String representation of the conditional node
func (c *Conditional) String() string {
	var out bytes.Buffer

	out.WriteString("if (")
	out.WriteString(c.condition.String())
	out.WriteString(") ")
	out.WriteString(c.thenStmt.String())

	if &c.elseStmt != nil {
		out.WriteString(" else ")
		out.WriteString(c.elseStmt.String())
	}

	return out.String()
}

// Build the symbol table for the conditional node
func (c *Conditional) BuildSymbolTable(tables *st.SymbolTables) {
	// Nothing to do here since any statements made are added to the symbol table at the function node level
}

// Type checking for the conditional node
func (c *Conditional) TypeCheck(errors []string, tables *st.SymbolTables) []string {
	return errors
}
