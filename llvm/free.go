package llvm

import (
	"bytes"
	"golite/arm"
	"golite/stack"
	"strconv"
)

// Representation of a free runtime call
type Free struct {
	blockLabel      string
	sourceRegisters []int
	targetRegisters []int
}

// NewFree returns a new free runtime call
func NewFree() *Free {
	srcR := make([]int, 0)
	// This is the most recent register used
	mostRecentR, _ := strconv.Atoi(GetPreviousRegister()[2:])
	srcR = append(srcR, mostRecentR)
	tgtR := make([]int, 0)
	return &Free{"", srcR, tgtR}
}

// String representation of the Local declaration
func (f *Free) String() string {
	var out bytes.Buffer
	// Format is - call void @free(i8* %reg_var)
	// Common for all
	out.WriteString("call void @free(i8* %")
	out.WriteString("r" + strconv.Itoa(f.sourceRegisters[len(f.sourceRegisters)-1]))
	out.WriteString(")")
	return out.String()
}

// Get the registers targeted by the instruction.
func (f *Free) GetTargets() []int {
	return f.targetRegisters
}

// Get the source registers of the instruction.
func (f *Free) GetSources() []int {
	return f.sourceRegisters
}

// Get the immediate value of the instruction.
func (f *Free) GetImmediate() int {
	return 0
}

// Get the label that marks this instruction in code.
func (f *Free) GetLabel() string {
	return f.blockLabel
}

// Set the label that marks this instruction in code.
func (f *Free) SetLabel(newLabel string) {
	f.blockLabel = newLabel
}

// Convert LLVM IR to ARM assembly.
func (f *Free) ToARM(stack *stack.Stack) []*arm.Instruction {
	return nil
}

// Build the stack table for the instruction.
func (f *Free) BuildStackTable(fName string, stack *stack.Stack) {
}
