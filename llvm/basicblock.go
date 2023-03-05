package llvm

import (
	"golite/arm"
	"golite/stack"
	"strconv"
)

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

// Get the predecessor labels.
func (bb *BasicBlock) GetPredecessorLabels() []string {
	labels := []string{}
	for _, pred := range bb.predecessors {
		labels = append(labels, pred.GetLabel())
	}
	return labels
}

// Get the successor labels.
func (bb *BasicBlock) GetSuccessorLabels() []string {
	labels := []string{}
	for _, succ := range bb.successors {
		labels = append(labels, succ.GetLabel())
	}
	return labels
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

// Conver the basic block to ARM assembly.
func (bb *BasicBlock) ToARM(funcName string, stack *stack.Stack, isFirstBlock bool) *arm.BasicBlock {
	armBB := arm.NewBasicBlock(bb.label)
	// Add the prologue to the first basic block.
	if isFirstBlock {
		// Grow the size of the stack by the size of the local variables.
		// amount = (total size of the stack - 1 (lr) - 1 (fp)) * 8 (size of a register)
		stackFrame := stack.GetFrame(funcName)
		growBy := (stackFrame.Size() - 2) * 8
		// Round to the nearest multiple of 16 for alignment.
		if growBy%16 != 0 {
			growBy += 16 - (growBy % 16)
		}

		if growBy > 0 {
			// Subtract SP by the amount.
			subInst := arm.NewSub(arm.SP, arm.SP, strconv.Itoa(growBy))
			subInst.SetLabel(armBB.GetLabel())
			armBB.AddInstruction(subInst)
		}

		// Subtract SP by 16 for the link register and the frame pointer.
		subInst := arm.NewSub(arm.SP, arm.SP, strconv.Itoa(16))
		subInst.SetLabel(armBB.GetLabel())
		armBB.AddInstruction(subInst)
		// Store the link register and the stack pointer.
		stpInst := arm.NewStp("x29", "x30", arm.SP)
		stpInst.SetLabel(armBB.GetLabel())
		armBB.AddInstruction(stpInst)
		// Move the stack pointer to the frame pointer.
		movInst := arm.NewMov(arm.FP, arm.SP)
		movInst.SetLabel(armBB.GetLabel())
		armBB.AddInstruction(movInst)

	}

	for _, inst := range bb.instructions {
		// It may be a 1:many mapping from LLVM IR to ARM assembly.
		insts := inst.ToARM(funcName, stack)
		for _, inst := range insts {
			armBB.AddInstruction(inst)
		}
	}

	return armBB
}

// Build the stack table
func (bb *BasicBlock) BuildStackTable(fnName string, stack *stack.Stack) {
	for _, inst := range bb.instructions {
		inst.BuildStackTable(fnName, stack)
	}
}
