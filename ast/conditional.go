package ast

import (
	"bytes"
	"golite/llvm"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
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
func (c *Conditional) BuildSymbolTable(tables *st.SymbolTables, errors []*SemanticAnalysisError) []*SemanticAnalysisError {
	// Nothing to do here since any statements made are added to the symbol table at the function node level
	return errors
}

// Type checking for the conditional node
func (c *Conditional) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables, funcEntry *st.FuncEntry) []*SemanticAnalysisError {
	// Type check the condition
	errors = c.condition.TypeCheck(errors, tables, funcEntry)

	// Type check the then block
	errors = c.thenBlock.TypeCheck(errors, tables, funcEntry)

	// Type check the else block
	if c.elseBlock != nil {
		errors = c.elseBlock.TypeCheck(errors, tables, funcEntry)
	}

	return errors
}

// Perform control flow analysis on the conditional node
func (c *Conditional) GetControlFlow(errors []*SemanticAnalysisError, funcEntry *st.FuncEntry) ([]*SemanticAnalysisError, bool) {
	// Perform control flow analysis on the then block
	var thenFlow bool
	errors, thenFlow = c.thenBlock.GetControlFlow(errors, funcEntry)

	// Perform control flow analysis on the else block
	var elseFlow bool
	if c.elseBlock != nil {
		errors, elseFlow = c.elseBlock.GetControlFlow(errors, funcEntry)
	}

	if funcEntry.RetTy != types.StringToType("void") {
		// If the else block does not exist, it doesnt matter if the then block has a return statement
		if c.elseBlock == nil {
			return errors, thenFlow
		} else {
			// If the then and else blocks exists, then both blocks must return true or both must return false
			if thenFlow == elseFlow {
				return errors, thenFlow
			} else {
				// Not adding the error below because it is possible that the parent node will have an appropriate return statement
				// errors = append(errors, NewSemanticAnalysisError("conditional branch missing control flow", "invalid control flow", c.Token))
				return errors, false
			}
		}
	} else {
		return errors, true
	}
}

// Translate the conditional node to LLVM IR
func (c *Conditional) ToLLVMCFG(tables *st.SymbolTables, blocks []*llvm.BasicBlock, funcEntry *st.FuncEntry) []*llvm.BasicBlock {

	return blocks
}
