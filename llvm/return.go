package llvm

import (
	"bytes"
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

	// struct type
	if strings.Contains(r.ty, "struct.") {
		out.WriteString("*")
	}

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
