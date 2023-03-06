package llvm

import (
	"bytes"
	"golite/arm"
	"golite/stack"
	"strconv"
	"strings"
)

// Representation of a binop operation
type BinOp struct {
	leftRegister    string
	op              string
	opType          string
	rightRegister   string
	blockLabel      string
	sourceRegisters []int
	targetRegisters []int
}

// NewBinOp returns a new free runtime call
func NewBinOp(leftRegister, op, opType, rightRegister string) *BinOp {
	srcR := make([]int, 0)
	// This is the most recent register used
	if strings.Contains(leftRegister, "%r") {
		mostRecentR, _ := strconv.Atoi(leftRegister[2:])
		srcR = append(srcR, mostRecentR)
	}

	// This always exists
	if strings.Contains(rightRegister, "%r") {
		mostRecentR, _ := strconv.Atoi(rightRegister[2:])
		srcR = append(srcR, mostRecentR)
	}

	// Get the next register
	tgtR := make([]int, 0)
	nextR, _ := strconv.Atoi(GetNextRegister()[2:])
	tgtR = append(tgtR, nextR)

	return &BinOp{leftRegister, op, opType, rightRegister, "", srcR, tgtR}
}

// String representation of the binop declaration
func (b *BinOp) String() string {
	var out bytes.Buffer
	// Format is - %result = <op> <type> %left, %right
	// <op> could be add, sub, etc., and, xor, etc., or icmp eq, icmp ne, etc.
	// Common for all
	out.WriteString("%r" + strconv.Itoa(b.targetRegisters[len(b.targetRegisters)-1]))
	out.WriteString(" = ")
	out.WriteString(b.op)
	out.WriteString(" ")
	if strings.Contains(b.opType, "struct.") {
		out.WriteString("%")
	}
	out.WriteString(b.opType)
	out.WriteString(" ")
	out.WriteString(b.leftRegister)
	out.WriteString(", ")
	out.WriteString(b.rightRegister)

	return out.String()
}

// Get the registers targeted by the instruction.
func (b *BinOp) GetTargets() []int {
	return b.targetRegisters
}

// Get the source registers of the instruction.
func (b *BinOp) GetSources() []int {
	return b.sourceRegisters
}

// Get the immediate value of the instruction.
func (b *BinOp) GetImmediate() int {
	return 0
}

// Get the label that marks this instruction in code.
func (b *BinOp) GetLabel() string {
	return b.blockLabel
}

// Set the label that marks this instruction in code.
func (b *BinOp) SetLabel(newLabel string) {
	b.blockLabel = newLabel
}

// Get the optype of the instruction.
func (b *BinOp) GetOpType() string {
	return b.opType
}

// Convert the BinOp to ARM assembly.
func (b *BinOp) ToARM(fnName string, stack *stack.Stack) []arm.Instruction {
	insts := make([]arm.Instruction, 0)

	// Left register could be a literal, global variable, address or register
	// literal value => move into a register
	// global variable => load from address into a register
	// an address => load from address into a register
	// a register => use the register

	// Right register could be a literal, global variable, address or register
	// literal value => move into a register
	// global variable => load from address into a register
	// an address => load from address into a register
	// a register => use the register

	// operator can be
	// add, sub, mul, sdiv
	// and, or, xor
	// eq, neq, sgt, sge, slt, sle
	return insts
}

// Function to build the stack table for the BinOp.
func (b *BinOp) BuildStackTable(funcName string, stack *stack.Stack) {
	destinationReg := "r" + strconv.Itoa(b.targetRegisters[len(b.targetRegisters)-1])
	stack.AddEntry(funcName, destinationReg, strconv.Itoa(stack.GetFrame(funcName).GetLargestOffset()+8))
}
