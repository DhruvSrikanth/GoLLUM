package arm

// Interface for instructions.
type Instruction interface {
	// Get the label that marks this instruction in code.
	GetLabel() string

	// Set the label that marks this instruction in code.
	SetLabel(newLabel string)

	// String Representation of the instruction.
	String() string
}
