package llvm

import (
	"bytes"
	"golite/arm"
	"golite/stack"
	"strconv"
	"strings"
)

// Representation of a bitcast operation
type BitCast struct {
	sourceRegister  string
	fromTy          string
	toTy            string
	blockLabel      string
	sourceRegisters []int
	targetRegisters []int
}

// NewBitCast returns a new BitCast operation
func NewBitCast(sourceRegister, fromTy, toTy string) *BitCast {
	srcR := make([]int, 0)
	// This is the most recent register used
	if strings.Contains(sourceRegister, "%r") {
		currReg, _ := strconv.Atoi(GetPreviousRegister()[2:])
		srcR = append(srcR, currReg)
	}
	tgtR := make([]int, 0)
	nextR, _ := strconv.Atoi(GetNextRegister()[2:])
	tgtR = append(tgtR, nextR)
	return &BitCast{sourceRegister, fromTy, toTy, "", srcR, tgtR}
}

// String representation of the Local declaration
func (b *BitCast) String() string {
	var out bytes.Buffer
	// Format is %result = bitcast type* %src_reg to type*
	// Common for all
	out.WriteString("%")
	out.WriteString("r" + strconv.Itoa(b.targetRegisters[len(b.targetRegisters)-1]))
	out.WriteString(" = bitcast ")
	// struct type
	if strings.Contains(b.fromTy, "struct.") {
		out.WriteString("%")
	}
	// Common for all
	out.WriteString(b.fromTy)
	out.WriteString(" ")
	out.WriteString(b.sourceRegister)
	out.WriteString(" to ")

	// struct type
	if strings.Contains(b.toTy, "struct.") {
		out.WriteString("%")
	}
	// Common for all
	out.WriteString(b.toTy)

	return out.String()
}

// Get the registers targeted by the instruction.
func (b *BitCast) GetTargets() []int {
	return b.targetRegisters
}

// Get the source registers of the instruction.
func (b *BitCast) GetSources() []int {
	return b.sourceRegisters
}

// Get the immediate value of the instruction.
func (b *BitCast) GetImmediate() int {
	return 0
}

// Get the label that marks this instruction in code.
func (b *BitCast) GetLabel() string {
	return b.blockLabel
}

// Set the label that marks this instruction in code.
func (b *BitCast) SetLabel(newLabel string) {
	b.blockLabel = newLabel
}

// Convert the instruction from LLVM IR to ARM assembly.
func (b *BitCast) ToARM(fnName string, stack *stack.Stack) []arm.Instruction {
	insts := make([]arm.Instruction, 0)

	stackFrame := stack.GetFrame(fnName)

	availableRegNum := stackFrame.GetNextRegister()
	availableReg := "x" + strconv.Itoa(availableRegNum)

	var srcR, destR string
	// The source could be a register or an address
	// an address => load from address into a register
	// a register => use the register
	srcAddr := stackFrame.GetLocation(b.sourceRegister[1:])
	if strings.Contains(srcAddr, "x") {
		// Register
		srcR = srcAddr
	} else {
		// Address

		ldrInst := arm.NewLdr(availableReg, arm.SP+", #"+srcAddr)
		ldrInst.SetLabel(b.blockLabel)
		insts = append(insts, ldrInst)

		srcR = availableReg
		availableRegNum += 1
	}

	// If the destination is a register, we can just move the value
	// else we will do a store in to the address
	destAddr := stackFrame.GetLocation("r" + strconv.Itoa(b.targetRegisters[len(b.targetRegisters)-1]))
	if strings.Contains(destAddr, "x") {
		// Register
		destR = destAddr
		movInst := arm.NewMov(destR, srcR)
		movInst.SetLabel(b.blockLabel)
		insts = append(insts, movInst)
	} else {
		// Address
		strInst := arm.NewStr(srcR, arm.SP+", #"+destAddr)
		strInst.SetLabel(b.blockLabel)
		insts = append(insts, strInst)
	}
	return insts
}

// Build the stack table for the instruction.
func (b *BitCast) BuildStackTable(funcName string, stack *stack.Stack) {
	destinationReg := "r" + strconv.Itoa(b.targetRegisters[len(b.targetRegisters)-1])
	stack.AddEntry(funcName, destinationReg, strconv.Itoa(stack.GetFrame(funcName).GetLargestOffset()+8))
}
