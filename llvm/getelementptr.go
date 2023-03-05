package llvm

import (
	"bytes"
	"golite/arm"
	"golite/stack"
	"strconv"
	"strings"
)

// Represents accessing a struct field
type GetElementPtr struct {
	sourceRegister  string
	ty              string
	index           int
	blockLabel      string
	sourceRegisters []int
	targetRegisters []int
}

// NewGetElementPtr returns a new getelementptr instruction
func NewGetElementPtr(sourceRegister string, ty string, index int) *GetElementPtr {
	srcR := make([]int, 0)
	if strings.Contains(sourceRegister, "%r") {
		sourceRInt, _ := strconv.Atoi(sourceRegister[2:])
		srcR = append(srcR, sourceRInt)
	}

	tgtR := make([]int, 0)
	nextR, _ := strconv.Atoi(GetNextRegister()[2:])
	tgtR = append(tgtR, nextR)

	return &GetElementPtr{sourceRegister, ty, index, "", srcR, tgtR}
}

// String representation of the getelementptr instruction
func (g *GetElementPtr) String() string {
	var out bytes.Buffer

	// Format is - %result = getelementptr ty, ty* %source, i32 0, i32 index
	// Common for all
	out.WriteString("%r" + strconv.Itoa(g.targetRegisters[len(g.targetRegisters)-1]))
	out.WriteString(" = getelementptr ")
	// struct type
	if strings.Contains(g.ty, "struct.") {
		out.WriteString("%")
	}
	out.WriteString(g.ty)
	out.WriteString(", ")
	// struct type
	if strings.Contains(g.ty, "struct.") {
		out.WriteString("%")
	}
	out.WriteString(g.ty)
	out.WriteString("* ")
	out.WriteString(g.sourceRegister)
	out.WriteString(", i32 0, i32 ")
	out.WriteString(strconv.Itoa(g.index))

	return out.String()
}

// Get the registers targeted by the instruction.
func (g *GetElementPtr) GetTargets() []int {
	return g.targetRegisters
}

// Get the source registers of the instruction.
func (g *GetElementPtr) GetSources() []int {
	return g.sourceRegisters
}

// Get the immediate value of the instruction.
func (g *GetElementPtr) GetImmediate() int {
	return 0
}

// Get the label that marks this instruction in code.
func (g *GetElementPtr) GetLabel() string {
	return g.blockLabel
}

// Set the label that marks this instruction in code.
func (g *GetElementPtr) SetLabel(newLabel string) {
	g.blockLabel = newLabel
}

// Convert from LLVM IR to ARM assembly.
func (g *GetElementPtr) ToARM() []*arm.Instruction {
	return nil
}

// Build the stack table for the function.
func (g *GetElementPtr) BuildStackTable(fName string, stack *stack.Stack) {
}
