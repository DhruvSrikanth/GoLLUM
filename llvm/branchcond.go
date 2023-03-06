package llvm

import (
	"bytes"
	"golite/arm"
	"golite/stack"
	"strconv"
	"strings"
)

// Represents an conditional branch instruction in LLVM IR
type BranchConditional struct {
	conditionReg    string
	trueLabel       string
	falseLabel      string
	blockLabel      string
	sourceRegisters []int
	targetRegisters []int
}

// NewBranchConditional creates a new branch instruction
func NewBranchConditional(conditionReg, trueLabel, falseLable string) *BranchConditional {
	srcR := make([]int, 0)
	if strings.Contains(conditionReg, "%r") {
		regNum, _ := strconv.Atoi(conditionReg[2:])
		srcR = append(srcR, regNum)
	}
	tgtR := make([]int, 0)
	return &BranchConditional{conditionReg, trueLabel, falseLable, "", srcR, tgtR}
}

// String representation of the branch instruction
func (b *BranchConditional) String() string {
	var out bytes.Buffer
	// Format is - br i1 %<condition_reg>, label <%true_label>, label <%false_label>
	out.WriteString("br i1 ")
	out.WriteString(b.conditionReg)
	out.WriteString(", label %")
	out.WriteString(b.trueLabel)
	out.WriteString(", label %")
	out.WriteString(b.falseLabel)
	return out.String()
}

// Get the true label
func (b *BranchConditional) GetTrueLabel() string {
	return b.trueLabel
}

// Get the false label
func (b *BranchConditional) GetFalseLabel() string {
	return b.falseLabel
}

// Get the condition register
func (b *BranchConditional) GetConditionRegister() string {
	return b.conditionReg
}

// Get the registers targeted by the instruction.
func (b *BranchConditional) GetTargets() []int {
	return b.targetRegisters
}

// Get the source registers of the instruction.
func (b *BranchConditional) GetSources() []int {
	return b.sourceRegisters
}

// Get the immediate value of the instruction.
func (b *BranchConditional) GetImmediate() int {
	return 0
}

// Get the label that marks this instruction in code.
func (b *BranchConditional) GetLabel() string {
	return b.blockLabel
}

// Set the label that marks this instruction in code.
func (b *BranchConditional) SetLabel(newLabel string) {
	b.blockLabel = newLabel
}

// Set the true label
func (b *BranchConditional) SetTrueLabel(newLabel string) {
	b.trueLabel = newLabel
}

// Set the false label
func (b *BranchConditional) SetFalseLabel(newLabel string) {
	b.falseLabel = newLabel
}

// Convert the instruction from LLVM IR to ARM assembly.
func (b *BranchConditional) ToARM(fnName string, stack *stack.Stack) []arm.Instruction {
	insts := make([]arm.Instruction, 0)

	return insts
}

// Build the stack table for the instruction.
func (b *BranchConditional) BuildStackTable(fName string, stack *stack.Stack) {
}
