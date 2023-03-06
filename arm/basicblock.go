package arm

// Basic Block representation for LLVM IR.
type BasicBlock struct {
	// The name of the basic block.
	label string
	// The instructions in the basic block.
	instructions []Instruction
}

// Create a new basic block.
func NewBasicBlock(label string) *BasicBlock {
	return &BasicBlock{
		label:        label,
		instructions: []Instruction{},
	}
}

// Get the label of the basic block.
func (bb *BasicBlock) GetLabel() string {
	return bb.label
}

// Get the instructions in the basic block.
func (bb *BasicBlock) GetInstructions() []Instruction {
	return bb.instructions
}

func (bb *BasicBlock) SetLabel(label string) {
	bb.label = label
}

// Set the instructions in the basic block.
func (bb *BasicBlock) SetInstructions(instructions []Instruction) {
	bb.instructions = instructions
}

// Add an instruction to the basic block.
func (bb *BasicBlock) AddInstruction(inst Instruction) {
	bb.instructions = append(bb.instructions, inst)
}

// Add multiple instructions to the basic block.
func (bb *BasicBlock) AddInstructions(instructions []Instruction) {
	bb.instructions = append(bb.instructions, instructions...)
}

// String representation of the basic block.
func (bb *BasicBlock) String() string {
	var out string
	out += "."
	out += bb.label
	out += ":\n"
	for _, inst := range bb.instructions {
		out += "\t"
		out += inst.String()
		out += "\n"
	}
	out += "\n"
	return out
}

// Get the last instruction in the basic block.
func (bb *BasicBlock) GetLastInstruction() Instruction {
	return bb.instructions[len(bb.instructions)-1]
}

// Get the number of instructions in the basic block.
func (bb *BasicBlock) Size() int {
	return len(bb.instructions)
}
