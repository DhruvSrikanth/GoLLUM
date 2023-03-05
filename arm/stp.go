package arm

// Store pair ARM64 instruction
type Stp struct {
	// The destination register
	Dest string
	// The first operand
	Op1 string
	// The second operand
	Op2 string

	blockLabel string
}

// NewStp returns a new Stp instruction
func NewStp(dest string, op1 string, op2 string) *Stp {
	return &Stp{dest, op1, op2, ""}
}

// String returns the string representation of the Stp instruction
func (s *Stp) String() string {
	return "stp " + s.Dest + ", " + s.Op1 + ", [" + s.Op2 + "]"
}

// Set the label of the instruction
func (s *Stp) SetLabel(label string) {
	s.blockLabel = label
}

// Get the label of the instruction
func (s *Stp) GetLabel() string {
	return s.blockLabel
}
