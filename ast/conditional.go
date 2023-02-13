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
	thenBlock *Block     // The block to be executed if the condition is true
	elseBlock *Block     // The block to be executed if the condition is false
}

// New Conditional node
func NewConditional(condition Expression, thenBlock *Block, elseBlock *Block, token *token.Token) *Conditional {
	return &Conditional{token, condition, thenBlock, elseBlock}
}

// String representation of the conditional node
func (c *Conditional) String() string {
	var out bytes.Buffer

	out.WriteString("if (")
	out.WriteString(c.condition.String())
	out.WriteString(") ")
	then := *c.thenBlock
	out.WriteString(then.String())

	if c.elseBlock != nil {
		out.WriteString(" else ")
		elseBlock := *c.elseBlock
		out.WriteString(elseBlock.String())
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
