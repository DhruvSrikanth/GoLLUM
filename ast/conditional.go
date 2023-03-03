package ast

import (
	"bytes"
	"golite/llvm"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
	"strconv"
	"strings"
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
func (c *Conditional) ToLLVMCFG(tables *st.SymbolTables, blocks []*llvm.BasicBlock, funcEntry *st.FuncEntry, constDecls []*llvm.ConstantDecl) ([]*llvm.BasicBlock, []*llvm.ConstantDecl) {
	var mostRecentOperand string

	// Before adding a new block, add an unconditional branch to the current block
	branchUncondInst := llvm.NewBranchUnconditional(llvm.GetCurrentLabel())
	branchUncondInst.SetLabel(blocks[len(blocks)-1].GetLabel())
	blocks[len(blocks)-1].AddInstruction(branchUncondInst)

	// Add new block for the condition
	condBlock := llvm.NewBasicBlock(llvm.GetNextLabel())
	// Add the last block to this block's predecessors and the previous block to this block's successors
	condBlock.AddPredecessor(blocks[len(blocks)-1])
	blocks[len(blocks)-1].AddSuccessor(condBlock)
	blocks = append(blocks, condBlock)

	// Evaluate condition expression to translate to LLVM IR
	blocks, constDecls, mostRecentOperand = c.condition.ToLLVMCFG(tables, blocks, funcEntry, constDecls)

	// If the last instruction is not a comparison i.e. the condition is a unary expression, we need to perform a comparison
	lastInst := blocks[len(blocks)-1].GetLastInstruction()
	if !strings.Contains(lastInst.String(), "icmp") {
		// Add a comparison instruction
		operationInst := llvm.NewBinOp(mostRecentOperand, llvm.OperatorToLLVM("=="), "i64", "1")
		operationInst.SetLabel(blocks[len(blocks)-1].GetLabel())
		blocks[len(blocks)-1].AddInstruction(operationInst)
		mostRecentOperand = llvm.GetPreviousRegister()
	}

	// Add a conditional branch instruction to the condition block
	condReg := mostRecentOperand                               // This is set in the condition expression
	currLabelID, _ := strconv.Atoi(llvm.GetCurrentLabel()[1:]) // remove the L from the label
	// These will have to change after the then and else blocks are added
	trueLabel := "L" + strconv.Itoa(currLabelID)
	falseLabel := "L" + strconv.Itoa(currLabelID+1)
	branchCondInst := llvm.NewBranchConditional(condReg, trueLabel, falseLabel)
	branchCondInst.SetLabel(blocks[len(blocks)-1].GetLabel())
	blocks[len(blocks)-1].AddInstruction(branchCondInst)

	// Add new block for the then block
	thenBlock := llvm.NewBasicBlock(llvm.GetNextLabel())
	// Add the last block to this block's predecessors and the previous block to this block's successors
	thenBlock.AddPredecessor(condBlock)
	condBlock.AddSuccessor(thenBlock)
	blocks = append(blocks, thenBlock)

	// Evaluate the then block before adding the unconditional branch instruction
	blocks, constDecls = c.thenBlock.ToLLVMCFG(tables, blocks, funcEntry, constDecls)

	// Unconditional branch to the block after the then and else blocks
	hasReturn := (blocks[len(blocks)-1].Size() > 0)
	if hasReturn {
		hasReturn = strings.Contains(blocks[len(blocks)-1].GetLastInstruction().String(), "ret ")
	}
	var branchUncondThenInst *llvm.BranchUnconditional
	if !hasReturn {
		branchUncondThenInst = llvm.NewBranchUnconditional(llvm.GetCurrentLabel())
		branchUncondThenInst.SetLabel(blocks[len(blocks)-1].GetLabel())
		blocks[len(blocks)-1].AddInstruction(branchUncondThenInst)
	}

	// Add new block for the else block
	elseBlock := llvm.NewBasicBlock(llvm.GetNextLabel())
	// Add the last block to this block's predecessors and
	elseBlock.AddPredecessor(condBlock)
	condBlock.AddSuccessor(elseBlock)
	blocks = append(blocks, elseBlock)

	// Update the conditions false label to the else block
	branchCondInst.SetFalseLabel(elseBlock.GetLabel())

	// Evaluate the else block if required before adding branch instructions
	if c.elseBlock != nil {
		blocks, constDecls = c.elseBlock.ToLLVMCFG(tables, blocks, funcEntry, constDecls)
	}

	// Update the then blocks unconditional branch instruction to the block after the then and else blocks if required
	if !hasReturn {
		branchUncondThenInst.SetDestinationLabel(llvm.GetCurrentLabel())
	}

	// Unconditional branch to the block after the then and else blocks
	hasReturn = (blocks[len(blocks)-1].Size() > 0)
	if hasReturn {
		hasReturn = strings.Contains(blocks[len(blocks)-1].GetLastInstruction().String(), "ret ")
	}
	if !hasReturn {
		branchUncondElseInst := llvm.NewBranchUnconditional(llvm.GetCurrentLabel())
		branchUncondElseInst.SetLabel(blocks[len(blocks)-1].GetLabel())
		blocks[len(blocks)-1].AddInstruction(branchUncondElseInst)
	}

	// Add new block for canonical form
	block := llvm.NewBasicBlock(llvm.GetNextLabel())
	// Add the last block to this block's predecessors
	block.AddPredecessor(thenBlock)
	block.AddPredecessor(elseBlock)
	blocks = append(blocks, block)

	return blocks, constDecls
}
