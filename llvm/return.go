package llvm

import (
	"bytes"
	"golite/arm"
	"golite/stack"
	"strconv"
	"strings"
)

// Represents a return instruction in LLVM IR
type Return struct {
	sourceRegister  string
	ty              string
	blockLabel      string
	sourceRegisters []int
	targetRegisters []int
}

// Create a new return instruction
func NewReturn(sourceRegister, ty string) *Return {
	srcR := make([]int, 0)
	targetR := make([]int, 0)
	return &Return{sourceRegister, ty, "", srcR, targetR}
}

// String representation of the return instruction
func (r *Return) String() string {
	// Output format is - ret <type> %<register>

	var out bytes.Buffer

	out.WriteString("ret ")

	// struct type
	if strings.Contains(r.ty, "struct.") {
		out.WriteString("%")
	}
	out.WriteString(r.ty)

	out.WriteString(" ")
	out.WriteString(r.sourceRegister)

	return out.String()
}

// Get the registers targeted by the instruction.
func (r *Return) GetTargets() []int {
	return r.targetRegisters
}

// Get the source registers of the instruction.
func (r *Return) GetSources() []int {
	return r.sourceRegisters
}

// Get the immediate value of the instruction.
func (r *Return) GetImmediate() int {
	return 0
}

// Get the label that marks this instruction in code.
func (r *Return) GetLabel() string {
	return r.blockLabel
}

// Set the label that marks this instruction in code.
func (r *Return) SetLabel(newLabel string) {
	r.blockLabel = newLabel
}

// Convert LLVM IR to ARM assembly.
func (r *Return) ToARM(fnName string, stack *stack.Stack) []arm.Instruction {
	instuctions := make([]arm.Instruction, 0)

	// Any time there is a return instruction, we need to have the epilogue
	// Restore the frame pointer and link register
	ldpInst := arm.NewLdp("x29", "x30", arm.SP)
	ldpInst.SetLabel(r.blockLabel)
	instuctions = append(instuctions, ldpInst)

	// Restore the stack pointer
	addInst := arm.NewAdd(arm.SP, arm.SP, "#16")
	addInst.SetLabel(r.blockLabel)
	instuctions = append(instuctions, addInst)

	// Add the amount that we grew the stack by previously
	stackFrame := stack.GetFrame(fnName)
	growBy := (stackFrame.Size() - 2) * 8
	// Round to the nearest multiple of 16 for alignment.
	if growBy%16 != 0 {
		growBy += 16 - (growBy % 16)
	}

	addInst = arm.NewAdd(arm.SP, arm.SP, "#"+strconv.Itoa(growBy))
	addInst.SetLabel(r.blockLabel)
	instuctions = append(instuctions, addInst)

	// Return
	retInst := arm.NewRet()
	retInst.SetLabel(r.blockLabel)
	instuctions = append(instuctions, retInst)

	return instuctions
}

// Build the stack table for the instruction.
func (r *Return) BuildStackTable(fName string, stack *stack.Stack) {
}
