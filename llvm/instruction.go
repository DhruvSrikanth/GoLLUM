package llvm

import (
	"golite/arm"
	"golite/stack"
)

// Interface for instructions.
type Instruction interface {
	// Get the registers targeted by the instruction.
	GetTargets() []int

	// Get the source registers of the instruction.
	GetSources() []int

	// Get the immediate value of the instruction.
	GetImmediate() int

	// Get the label that marks this instruction in code.
	GetLabel() string

	// Set the label that marks this instruction in code.
	SetLabel(newLabel string)

	// String Representation of the instruction.
	String() string

	// Convert the instruction to ARM assembly.
	ToARM(*stack.Stack) []*arm.Instruction

	// Build the stack table for the instruction.
	BuildStackTable(funcName string, stack *stack.Stack)
}
