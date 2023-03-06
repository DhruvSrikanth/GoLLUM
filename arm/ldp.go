package arm

// Load pair ARM64 instruction
type Ldp struct {
	// The destination register
	Dest string
	// The first operand
	Op1 string
	// The second operand
	Op2 string

	blockLabel string
}

// NewLdp returns a new Ldp instruction
func NewLdp(dest string, op1 string, op2 string) *Ldp {
	return &Ldp{dest, op1, op2, ""}
}

// String returns the string representation of the Ldp instruction
func (l *Ldp) String() string {
	return "ldp " + l.Dest + ", " + l.Op1 + ", [" + l.Op2 + "]"
}

// Set the label of the instruction
func (l *Ldp) SetLabel(label string) {
	l.blockLabel = label
}

// Get the label of the instruction
func (l *Ldp) GetLabel() string {
	return l.blockLabel
}
