package arm

// Mov ARM64 instruction
type Mov struct {
	// The destination register
	Dest string
	// The operand
	Op string
	// block label
	blockLabel string
}

// NewMov returns a new Mov instruction
func NewMov(dest string, op string) *Mov {
	return &Mov{dest, op, ""}
}

// String returns the string representation of the Mov instruction
func (m *Mov) String() string {
	return "mov " + m.Dest + ", " + m.Op
}

// Set the label of the instruction
func (m *Mov) SetLabel(label string) {
	m.blockLabel = label
}

// Get the label of the instruction
func (m *Mov) GetLabel() string {
	return m.blockLabel
}
