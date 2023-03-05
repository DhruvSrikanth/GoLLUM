package arm

// LDR ARM64 instruction
type Ldr struct {
	// The target register
	target string
	// The source address
	sourceAddress string

	// block label
	blockLabel string
}

// NewLDR returns a new LDR instruction
func NewLdr(target string, sourceAddress string) *Ldr {
	return &Ldr{target, sourceAddress, ""}
}

// String returns the string representation of the LDR instruction
func (l *Ldr) String() string {
	return "ldr " + l.target + ", [" + l.sourceAddress + "]"
}

// Set the label of the instruction
func (l *Ldr) SetLabel(label string) {
	l.blockLabel = label
}

// Get the label of the instruction
func (l *Ldr) GetLabel() string {
	return l.blockLabel
}
