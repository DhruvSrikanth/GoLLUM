package arm

// Return ARM64 instruction
type Ret struct {
	blockLabel string
}

// NewRet returns a new Ret instruction
func NewRet() *Ret {
	return &Ret{}
}

// String returns the string representation of the return instruction
func (r *Ret) String() string {
	return "ret"
}

// Set the label of the instruction
func (r *Ret) SetLabel(label string) {
	r.blockLabel = label
}

// Get the label of the instruction
func (r *Ret) GetLabel() string {
	return r.blockLabel
}
