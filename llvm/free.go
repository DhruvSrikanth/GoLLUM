package llvm

import (
	"bytes"
	"golite/arm"
	"golite/stack"
	"strconv"
	"strings"
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
func (f *Free) ToARM(fnName string, stack *stack.Stack) []arm.Instruction {
	insts := make([]arm.Instruction, 0)

	stackFrame := stack.GetFrame(fnName)
	availableRegNum := stackFrame.GetNextRegister()
	availableReg := "x" + strconv.Itoa(availableRegNum)

	// move whatever is in x0 to the next available register
	movInst := arm.NewMov(availableReg, "x0")
	availableRegNum += 1
	movInst.SetLabel(f.blockLabel)
	insts = append(insts, movInst)

	// Get the address of the register that holds the value to be freed
	sourceReg := "r" + strconv.Itoa(f.sourceRegisters[len(f.sourceRegisters)-1])
	sourceRegAddress := stackFrame.GetLocation(sourceReg)

	if strings.Contains(sourceRegAddress, "x") {
		// Store the result in the stack by doing a mov
		movInst := arm.NewMov("x0", sourceRegAddress)
		movInst.SetLabel(f.blockLabel)
		insts = append(insts, movInst)
	} else {
		// Load the value from the stack
		availableReg = "x" + strconv.Itoa(availableRegNum)
		ldrInst := arm.NewLdr(availableReg, sourceRegAddress)
		availableRegNum += 1
		ldrInst.SetLabel(f.blockLabel)
		insts = append(insts, ldrInst)

		// Now move the value to x0
		movInst := arm.NewMov("x0", availableReg)
		availableRegNum -= 1
		movInst.SetLabel(f.blockLabel)
		insts = append(insts, movInst)

	}
	availableRegNum -= 1

	// Call the free runtime function
	callInst := arm.NewBranch("free")
	callInst.SetLabel(f.blockLabel)
	insts = append(insts, callInst)

	// Restore the value of x0
	availableReg = "x" + strconv.Itoa(availableRegNum)
	movInst = arm.NewMov("x0", availableReg)
	movInst.SetLabel(f.blockLabel)
	insts = append(insts, movInst)

	return insts
}

// Build the stack table for the instruction.
func (f *Free) BuildStackTable(fName string, stack *stack.Stack) {
}
