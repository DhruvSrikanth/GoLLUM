package arm

// Movk ARM64 instruction
type Movk struct {
	// The destination register
	dest   string
	source string
	// block label
	blockLabel string
}

// NewMovk returns a new Mov instruction
func NewMovk(dest string, src string) *Movk {
	return &Movk{dest, src, ""}
}

// String returns the string representation of the Movk instruction
func (m *Movk) String() string {
	return "movk " + m.dest + ", " + m.source
}

// Set the label of the instruction
func (m *Movk) SetLabel(label string) {
	m.blockLabel = label
}

// Get the label of the instruction
func (m *Movk) GetLabel() string {
	return m.blockLabel
}
