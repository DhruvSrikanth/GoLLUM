package llvm

import "bytes"

// Represents an unconditional branch instruction in LLVM IR
type BranchUnconditional struct {
	destinationLabel string // The label to branch to
	blockLabel       string
	sourceRegisters  []int
	targetRegisters  []int
}

// NewBranchUnconditional creates a new branch instruction
func NewBranchUnconditional(destinationLabel string) *BranchUnconditional {
	srcR := make([]int, 0)
	tgtR := make([]int, 0)
	return &BranchUnconditional{destinationLabel, "", srcR, tgtR}
}

// String representation of the branch instruction
func (b *BranchUnconditional) String() string {
	var out bytes.Buffer
	// Format is - br label <%label>
	out.WriteString("br label %")
	out.WriteString(b.destinationLabel)
	return out.String()
}

// Get the destination label
func (b *BranchUnconditional) GetDestinationLabel() string {
	return b.destinationLabel
}

// Get the registers targeted by the instruction.
func (b *BranchUnconditional) GetTargets() []int {
	return b.targetRegisters
}

// Get the source registers of the instruction.
func (b *BranchUnconditional) GetSources() []int {
	return b.sourceRegisters
}

// Get the immediate value of the instruction.
func (b *BranchUnconditional) GetImmediate() int {
	return 0
}

// Get the label that marks this instruction in code.
func (b *BranchUnconditional) GetLabel() string {
	return b.blockLabel
}

// Set the label that marks this instruction in code.
func (b *BranchUnconditional) SetLabel(newLabel string) {
	b.blockLabel = newLabel
}
