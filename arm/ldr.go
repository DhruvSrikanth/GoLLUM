package arm

// LDR ARM64 instruction
type Ldr struct {
	// The target register
	Target string
	// The offset of the source register from sp
	SPOffset string

	// block label
	blockLabel string
}

// NewLDR returns a new LDR instruction
func NewLdr(target string, spOffset string) *Ldr {
	return &Ldr{target, spOffset, ""}
}

// String returns the string representation of the LDR instruction
func (l *Ldr) String() string {
	return "ldr " + l.Target + ", [sp, #" + l.SPOffset + "]"
}

// Set the label of the instruction
func (l *Ldr) SetLabel(label string) {
	l.blockLabel = label
}

// Get the label of the instruction
func (l *Ldr) GetLabel() string {
	return l.blockLabel
}
