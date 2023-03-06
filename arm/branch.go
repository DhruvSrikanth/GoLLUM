package arm

import "fmt"

// Branch to label ARM64 instruction
type Branch struct {
	// The label to branch to
	label      string
	blockLabel string
	branchType string
}

// NewBranch returns a new branch instruction
func NewBranch(label string) *Branch {
	return &Branch{label, "", "function"}
}

// String representation of the branch instruction
func (b *Branch) String() string {
	if b.branchType == "function" {
		return "bl " + b.label
	} else if b.branchType == "cfg" {
		return "b " + b.label
	} else {
		fmt.Println("[error in branch translation]")
		return ""
	}
}

// Get the block label of the instruction.
func (b *Branch) GetLabel() string {
	return b.blockLabel
}

// Set the block label of the instruction.
func (b *Branch) SetLabel(newLabel string) {
	b.blockLabel = newLabel
}

// Get the branch type of the instruction.
func (b *Branch) GetBranchType() string {
	return b.branchType
}

// Set the branch type of the instruction.
func (b *Branch) SetBranchType(newType string) {
	b.branchType = newType
}
