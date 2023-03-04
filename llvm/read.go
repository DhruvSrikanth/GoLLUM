package llvm

import (
	"bytes"
	"golite/arm"
	"strconv"
	"strings"
)

// Representation of a Read runtime call
type Read struct {
	sourceReg       string
	blockLabel      string
	sourceRegisters []int
	targetRegisters []int
}

// NewRead returns a new Read runtime call
func NewRead(sourceRegister string) *Read {
	srcR := make([]int, 0)
	if strings.Contains(sourceRegister, "%r") {
		rInt, _ := strconv.Atoi(sourceRegister[2:])
		srcR = append(srcR, rInt)
	}

	tgtR := make([]int, 0)

	return &Read{
		sourceReg:       sourceRegister,
		blockLabel:      "",
		sourceRegisters: srcR,
		targetRegisters: tgtR,
	}
}

// String representation of the Read call
func (r *Read) String() string {
	var out bytes.Buffer
	// Format is -  call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %register)
	// Common for all
	out.WriteString("call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* ")
	out.WriteString(r.sourceReg)
	out.WriteString(")")

	return out.String()
}

// Get the registers targeted by the instruction.
func (r *Read) GetTargets() []int {
	return r.targetRegisters
}

// Get the source registers of the instruction.
func (r *Read) GetSources() []int {
	return r.sourceRegisters
}

// Get the immediate value of the instruction.
func (r *Read) GetImmediate() int {
	return 0
}

// Get the label that marks this instruction in code.
func (r *Read) GetLabel() string {
	return r.blockLabel
}

// Set the label that marks this instruction in code.
func (r *Read) SetLabel(newLabel string) {
	r.blockLabel = newLabel
}

// Convert LLVM IR to ARM assembly.
func (r *Read) ToARM() []*arm.Instruction {
	return nil
}
