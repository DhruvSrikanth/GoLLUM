package arm

// Movz ARM64 instruction
type Movz struct {
	// The destination register
	dest   string
	source string
	// block label
	blockLabel string
}

// NewMovz returns a new Mov instruction
func NewMovz(dest string, src string) *Movz {
	return &Movz{dest, src, ""}
}

// String returns the string representation of the Movz instruction
func (m *Movz) String() string {
	return "movz " + m.dest + ", " + m.source
}

// Set the label of the instruction
func (m *Movz) SetLabel(label string) {
	m.blockLabel = label
}

// Get the label of the instruction
func (m *Movz) GetLabel() string {
	return m.blockLabel
}
