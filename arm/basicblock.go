package arm

// Basic Block representation for LLVM IR.
type BasicBlock struct {
	// The name of the basic block.
	label string
	// The instructions in the basic block.
	instructions []Instruction
	// The predecessor and successor basic blocks.
	predecessors []*BasicBlock
	successors   []*BasicBlock
}

// Create a new basic block.
func NewBasicBlock(label string) *BasicBlock {
	return &BasicBlock{
		label:        label,
		instructions: []Instruction{},
		predecessors: []*BasicBlock{},
		successors:   []*BasicBlock{},
	}
}

// Add a predecessor basic block.
func (bb *BasicBlock) AddPredecessor(pred *BasicBlock) {
	bb.predecessors = append(bb.predecessors, pred)
}

// Add a successor basic block.
func (bb *BasicBlock) AddSuccessor(succ *BasicBlock) {
	bb.successors = append(bb.successors, succ)
}

// Get the label of the basic block.
func (bb *BasicBlock) GetLabel() string {
	return bb.label
}

// Get the instructions in the basic block.
func (bb *BasicBlock) GetInstructions() []Instruction {
	return bb.instructions
}

// Get the predecessor basic blocks.
func (bb *BasicBlock) GetPredecessors() []*BasicBlock {
	return bb.predecessors
}

// Get the successor basic blocks.
func (bb *BasicBlock) GetSuccessors() []*BasicBlock {
	return bb.successors
}

// Set the label of the basic block.
func (bb *BasicBlock) SetLabel(label string) {
	bb.label = label
}

// Set the instructions in the basic block.
func (bb *BasicBlock) SetInstructions(instructions []Instruction) {
	bb.instructions = instructions
}

// Set the predecessor basic blocks.
func (bb *BasicBlock) SetPredecessors(predecessors []*BasicBlock) {
	bb.predecessors = predecessors
}

// Set the successor basic blocks.
func (bb *BasicBlock) SetSuccessors(successors []*BasicBlock) {
	bb.successors = successors
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