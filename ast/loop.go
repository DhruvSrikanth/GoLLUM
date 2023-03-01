package ast

import (
	"bytes"
	"golite/llvm"
	st "golite/symboltable"
	"golite/token"
	"strconv"
)

// Loop node for the AST
type Loop struct {
	*token.Token
	condition Expression // The condition for the loop
	body      Block      // The block to be executed if the condition is true
}

// NewLoop node
func NewLoop(condition Expression, body Block, token *token.Token) *Loop {
	return &Loop{token, condition, body}
}

// String representation of the loop node
func (l *Loop) String() string {
	var out bytes.Buffer

	out.WriteString("for (")
	out.WriteString(l.condition.String())
	out.WriteString(") ")
	out.WriteString(l.body.String())

	return out.String()
}

// Build the symbol table for the loop node
func (l *Loop) BuildSymbolTable(tables *st.SymbolTables, errors []*SemanticAnalysisError) []*SemanticAnalysisError {
	// Nothing to do here since the body is already in the symbol table
	return errors
}

// Type checking for the loop node
func (l *Loop) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables, funcEntry *st.FuncEntry) []*SemanticAnalysisError {
	// Type check the condition
	errors = l.condition.TypeCheck(errors, tables, funcEntry)

	// Type check the body
	errors = l.body.TypeCheck(errors, tables, funcEntry)

	return errors
}

// Control flow analysis for the loop node
func (l *Loop) GetControlFlow(errors []*SemanticAnalysisError, funcEntry *st.FuncEntry) ([]*SemanticAnalysisError, bool) {
	// Perform control flow analysis on the block
	var flow bool
	errors, flow = l.body.GetControlFlow(errors, funcEntry)
	return errors, flow
}

// Translate the loop node to LLVM IR
func (l *Loop) ToLLVMCFG(tables *st.SymbolTables, blocks []*llvm.BasicBlock, funcEntry *st.FuncEntry, constDecls []*llvm.ConstantDecl) ([]*llvm.BasicBlock, []*llvm.ConstantDecl) {
	var mostRecentOperand string
	// Before adding a new block, add an unconditional branch to the condition block
	branchUncondInst := llvm.NewBranchUnconditional(llvm.GetCurrentLabel())
	branchUncondInst.SetLabel(blocks[len(blocks)-1].GetLabel())
	blocks[len(blocks)-1].AddInstruction(branchUncondInst)

	// Add new block for the condition
	condBlock := llvm.NewBasicBlock(llvm.GetNextLabel())
	// Add the previous block to this blocks predecessors
	condBlock.AddPredecessor(blocks[len(blocks)-1])
	// Add this block to the last blocks successors
	blocks[len(blocks)-1].AddSuccessor(condBlock)
	// Add the condition instructions to the block
	blocks = append(blocks, condBlock)
	// Get the condition expression to translate to LLVM IR
	blocks, constDecls, mostRecentOperand = l.condition.ToLLVMCFG(tables, blocks, funcEntry, constDecls)

	// Add a conditional branch to the body
	condReg := mostRecentOperand
	currLabelID, _ := strconv.Atoi(llvm.GetCurrentLabel()[1:]) // remove the L from the label
	trueLabel := "L" + strconv.Itoa(currLabelID)
	falseLabel := "L" + strconv.Itoa(currLabelID+1)

	branchCondInst := llvm.NewBranchConditional(condReg, trueLabel, falseLabel)
	branchCondInst.SetLabel(blocks[len(blocks)-1].GetLabel())
	blocks[len(blocks)-1].AddInstruction(branchCondInst)

	// Add new block for the body block
	bodyBlock := llvm.NewBasicBlock(llvm.GetNextLabel())
	// Add the previous block to this blocks predecessors
	bodyBlock.AddPredecessor(blocks[len(blocks)-1])
	// Add this block to the last blocks successors
	blocks[len(blocks)-1].AddSuccessor(bodyBlock)

	// Append the then block to the list of blocks since the last block will be considered for the block to add instructions to
	blocks = append(blocks, bodyBlock)
	blocks, constDecls = l.body.ToLLVMCFG(tables, blocks, funcEntry, constDecls)

	// Add an unconditional branch to the condition block
	condLabelID := currLabelID - 1
	condLabel := "L" + strconv.Itoa(condLabelID)
	branchUncondInst = llvm.NewBranchUnconditional(condLabel)
	branchUncondInst.SetLabel(blocks[len(blocks)-1].GetLabel())
	blocks[len(blocks)-1].AddInstruction(branchUncondInst)

	// Add new block for canonical form
	block := llvm.NewBasicBlock(llvm.GetNextLabel())
	// Add the previous block to this blocks predecessors
	block.AddPredecessor(blocks[len(blocks)-1])
	// Add the cond block to this blocks predecessors
	block.AddPredecessor(blocks[len(blocks)-2])
	// Add this block to the last blocks successors
	blocks[len(blocks)-1].AddSuccessor(block)
	blocks[len(blocks)-2].AddSuccessor(block)
	// Add the block to the list of blocks
	blocks = append(blocks, block)

	return blocks, constDecls
}
