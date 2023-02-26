package llvm

import (
	"bytes"
	"strconv"
	"strings"
)

// Representation of a bitcast operation
type BitCast struct {
	sourceRegister  string
	fromTy          string
	toTy            string
	blockLabel      string
	sourceRegisters []int
	targetRegisters []int
}

// NewBitCast returns a new BitCast operation
func NewBitCast(sourceRegister, fromTy, toTy string) *BitCast {
	srcR := make([]int, 0)
	// This is the most recent register used
	if strings.Contains(sourceRegister, "%r") {
		currReg, _ := strconv.Atoi(GetPreviousRegister()[2:])
		srcR = append(srcR, currReg)
	}
	tgtR := make([]int, 0)
	nextR, _ := strconv.Atoi(GetNextRegister()[2:])
	tgtR = append(tgtR, nextR)
	return &BitCast{sourceRegister, fromTy, toTy, "", srcR, tgtR}
}

// String representation of the Local declaration
func (b *BitCast) String() string {
	var out bytes.Buffer
	// Format is %result = bitcast type* %src_reg to type*
	// Common for all
	out.WriteString("%")
	out.WriteString("r" + strconv.Itoa(b.targetRegisters[len(b.targetRegisters)-1]))
	out.WriteString(" = bitcast ")
	// struct type
	if strings.Contains(b.fromTy, "struct.") {
		out.WriteString("%")
	}
	// Common for all
	out.WriteString(b.fromTy)
	out.WriteString("* ")
	out.WriteString(b.sourceRegister)
	out.WriteString(" to ")

	// struct type
	if strings.Contains(b.toTy, "struct.") {
		out.WriteString("%")
	}
	// Common for all
	out.WriteString(b.toTy)
	out.WriteString("*")

	return out.String()
}

// Get the registers targeted by the instruction.
func (b *BitCast) GetTargets() []int {
	return b.targetRegisters
}

// Get the source registers of the instruction.
func (b *BitCast) GetSources() []int {
	return b.sourceRegisters
}

// Get the immediate value of the instruction.
func (b *BitCast) GetImmediate() int {
	return 0
}

// Get the label that marks this instruction in code.
func (b *BitCast) GetLabel() string {
	return b.blockLabel
}

// Set the label that marks this instruction in code.
func (b *BitCast) SetLabel(newLabel string) {
	b.blockLabel = newLabel
}
