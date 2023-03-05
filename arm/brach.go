package arm

// Branch to label ARM64 instruction
type Branch struct {
	// The label to branch to
	label string

	blockLabel string
}

// NewBranch returns a new branch instruction
func NewBranch(label string) *Branch {
	return &Branch{label, ""}
}

// String representation of the branch instruction
func (b *Branch) String() string {
	return "bl " + b.label
}

// Get the block label of the instruction.
func (b *Branch) GetLabel() string {
	return b.blockLabel
}

// Set the block label of the instruction.
func (b *Branch) SetLabel(newLabel string) {
	b.blockLabel = newLabel
}
