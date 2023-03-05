package arm

// Mov ARM64 instruction
type Mov struct {
	// The destination register
	Dest   string
	Source string
	// block label
	blockLabel string
}

// NewMov returns a new Mov instruction
func NewMov(dest string, src string) *Mov {
	return &Mov{dest, src, ""}
}

// String returns the string representation of the Mov instruction
func (m *Mov) String() string {
	return "mov " + m.Dest + ", " + m.Source
}

// Set the label of the instruction
func (m *Mov) SetLabel(label string) {
	m.blockLabel = label
}

// Get the label of the instruction
func (m *Mov) GetLabel() string {
	return m.blockLabel
}
