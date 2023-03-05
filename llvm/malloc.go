package llvm

import (
	"bytes"
	"golite/arm"
	"golite/stack"
	"strconv"
)

// Representation of a malloc runtime call
type Malloc struct {
	size            int
	blockLabel      string
	sourceRegisters []int
	targetRegisters []int
}

// Newmalloc returns a new malloc runtime call
func NewMalloc(size int) *Malloc {
	srcR := make([]int, 0)
	tgtR := make([]int, 0)
	nextR, _ := strconv.Atoi(GetNextRegister()[2:])
	tgtR = append(tgtR, nextR)
	return &Malloc{size, "", srcR, tgtR}
}

// String representation of the malloc call
func (m *Malloc) String() string {
	var out bytes.Buffer
	// Format is - %result = call i8* @malloc(i32 size)
	// Common for all
	out.WriteString("%r" + strconv.Itoa(m.targetRegisters[len(m.targetRegisters)-1]))
	out.WriteString(" = call i8* @malloc(i32 ")
	out.WriteString(strconv.Itoa(m.size))
	out.WriteString(")")
	return out.String()
}

// Get the registers targeted by the instruction.
func (m *Malloc) GetTargets() []int {
	return m.targetRegisters
}

// Get the source registers of the instruction.
func (m *Malloc) GetSources() []int {
	return m.sourceRegisters
}

// Get the immediate value of the instruction.
func (m *Malloc) GetImmediate() int {
	return 0
}

// Get the label that marks this instruction in code.
func (m *Malloc) GetLabel() string {
	return m.blockLabel
}

// Set the label that marks this instruction in code.
func (m *Malloc) SetLabel(newLabel string) {
	m.blockLabel = newLabel
}

// Convert LLVM IR to ARM assembly.
func (m *Malloc) ToARM(stack *stack.Stack) []*arm.Instruction {
	return nil
}

// Build the stack table for the instruction.
func (m *Malloc) BuildStackTable(funcName string, stack *stack.Stack) {
	destinationReg := "r" + strconv.Itoa(m.targetRegisters[len(m.targetRegisters)-1])
	stack.AddEntry(funcName, destinationReg, strconv.Itoa(stack.GetFrame(funcName).GetLargestOffset()+8))
}
