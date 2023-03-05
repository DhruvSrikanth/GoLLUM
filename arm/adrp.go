package arm

// ADRP ARM64 instruction
type Adrp struct {
	// The name of the register to store the result in
	Dest string
	// The label to calculate the address of
	Label string

	blockLabel string
}

// NewAdrp returns a new ADRP instruction
func NewAdrp(dest, label string) *Adrp {
	return &Adrp{dest, label, ""}
}

// String representation of the ADRP instruction
func (a *Adrp) String() string {
	var out string
	out += "adrp " + a.Dest + ", " + a.Label
	return out
}

// Set the label of the instruction
func (a *Adrp) SetLabel(label string) {
	a.blockLabel = label
}

// Get the label of the instruction
func (a *Adrp) GetLabel() string {
	return a.blockLabel
}
