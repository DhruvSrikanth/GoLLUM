package llvm

import "golite/arm"

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
	ToARM() []*arm.Instruction
}
