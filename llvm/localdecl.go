package llvm

import (
	"bytes"
	"golite/arm"
	"golite/stack"
	"strconv"
	"strings"
)

// Representation of a Local declaration
type LocalDecl struct {
	// The name of the Local declaration
	name       string
	ty         string
	blockLabel string
}

// NewLocalDecl returns a new Local declaration
func NewLocalDecl(name, ty string) *LocalDecl {
	return &LocalDecl{name, ty, ""}
}

// String representation of the Local declaration
func (s *LocalDecl) String() string {
	var out bytes.Buffer
	// Format is %var_name = alloca type
	// Common for all
	out.WriteString("%")
	out.WriteString(s.name)
	out.WriteString(" = alloca ")
	// struct type
	if strings.Contains(s.ty, "struct.") {
		out.WriteString("%")
	}
	// Common for all
	out.WriteString(s.ty)
	if strings.Contains(s.ty, "struct.") {
		out.WriteString("*")
	}
	return out.String()
}

// Get the registers targeted by the instruction.
func (s *LocalDecl) GetTargets() []int {
	return []int{}
}

// Get the source registers of the instruction.
func (s *LocalDecl) GetSources() []int {
	return []int{}
}

// Get the immediate value of the instruction.
func (s *LocalDecl) GetImmediate() int {
	return 0
}

// Get the label that marks this instruction in code.
func (s *LocalDecl) GetLabel() string {
	return s.blockLabel
}

// Set the label that marks this instruction in code.
func (s *LocalDecl) SetLabel(newLabel string) {
	s.blockLabel = newLabel
}

// Convert LLVM IR to ARM assembly.
func (s *LocalDecl) ToARM(fnName string, stack *stack.Stack) []arm.Instruction {
	// Nothing to return since we allocate the required amount of space in the prologue which is at the start of the function
	return []arm.Instruction{}
}

// Build the stack table for the instruction.
func (s *LocalDecl) BuildStackTable(fName string, stack *stack.Stack) {
	if !strings.Contains(s.name, "P_") {
		stack.AddEntry(fName, s.name, strconv.Itoa(stack.GetFrame(fName).GetLargestOffset()+8))
	}
}
