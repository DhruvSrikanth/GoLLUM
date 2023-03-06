package llvm

import (
	"bytes"
	"golite/arm"
	"golite/stack"
	"strconv"
	"strings"
)

// Representation of a load operation
type Load struct {
	sourceRegister  string
	ty              string
	blockLabel      string
	sourceRegisters []int
	targetRegisters []int
}

// NewLoad returns a new Load operation
func NewLoad(sourceRegister, ty string) *Load {
	srcR := make([]int, 0)
	if strings.Contains(sourceRegister, "%r") {
		mostRecentR, _ := strconv.Atoi(GetPreviousRegister()[2:])
		srcR = append(srcR, mostRecentR)
	}
	tgtR := make([]int, 0)
	nextR, _ := strconv.Atoi(GetNextRegister()[2:])
	tgtR = append(tgtR, nextR)
	return &Load{sourceRegister, ty, "", srcR, tgtR}
}

// String representation of the Local declaration
func (l *Load) String() string {
	var out bytes.Buffer
	// Format is %result = load type, type* %var_name
	// Common for all
	out.WriteString("%")
	out.WriteString("r" + strconv.Itoa(l.targetRegisters[len(l.targetRegisters)-1]))
	out.WriteString(" = load ")
	// struct type
	if strings.Contains(l.ty, "struct.") {
		out.WriteString("%")
	}
	// Common for all
	out.WriteString(l.ty)

	// Common for all
	out.WriteString(", ")
	// struct type
	if strings.Contains(l.ty, "struct.") {
		out.WriteString("%")
	}
	// Common for all
	out.WriteString(l.ty)
	out.WriteString("* ")
	out.WriteString(l.sourceRegister)

	return out.String()
}

// Get the registers targeted by the instruction.
func (l *Load) GetTargets() []int {
	return l.targetRegisters
}

// Get the source registers of the instruction.
func (l *Load) GetSources() []int {
	return l.sourceRegisters
}

// Get the immediate value of the instruction.
func (l *Load) GetImmediate() int {
	return 0
}

// Get the label that marks this instruction in code.
func (l *Load) GetLabel() string {
	return l.blockLabel
}

// Set the label that marks this instruction in code.
func (l *Load) SetLabel(newLabel string) {
	l.blockLabel = newLabel
}

// Convert LLVM IR to ARM assembly.
func (l *Load) ToARM(fnName string, stack *stack.Stack) []arm.Instruction {
	insts := make([]arm.Instruction, 0)
	stackFrame := stack.GetFrame(fnName)
	availableRegNum := stackFrame.GetNextRegister()
	availableReg := "x" + strconv.Itoa(availableRegNum)

	var toR, fromR string

	// Source registers/addresses
	if strings.Contains(l.sourceRegister, "@") {
		// If its a global variable, then we need to load the address in a different way
		adrpInst := arm.NewAdrp(availableReg, l.sourceRegister[1:])
		adrpInst.SetLabel(l.blockLabel)
		insts = append(insts, adrpInst)

		addInst := arm.NewAdd(availableReg, availableReg, ":lo12:"+l.sourceRegister[1:])
		addInst.SetLabel(l.blockLabel)
		insts = append(insts, addInst)

		// Load the value into the register
		ldrInst := arm.NewLdr(availableReg, availableReg)
		ldrInst.SetLabel(l.blockLabel)
		insts = append(insts, ldrInst)

		availableRegNum += 1

		fromR = availableReg
	} else {
		src := l.sourceRegister[1:]
		destAddr := stackFrame.GetLocation(src)
		if strings.Contains(destAddr, "x") {
			fromR = destAddr
		} else {
			// Load the address of the variable into a register
			ldrInst := arm.NewLdr(availableReg, arm.SP+", #"+destAddr)
			ldrInst.SetLabel(l.blockLabel)
			insts = append(insts, ldrInst)

			availableRegNum += 1

			fromR = availableReg
		}
	}

	// Destination registers/addresses (it wont be a global variable since we are loading into a new register i.e. in LLVM IR will have %r for static single assignment)
	dest := "r" + strconv.Itoa(l.targetRegisters[len(l.targetRegisters)-1])
	destAddr := stackFrame.GetLocation(dest)
	if strings.Contains(destAddr, "x") {
		// Do a move since they are both in registers
		toR = destAddr

		movInst := arm.NewMov(toR, fromR)
		movInst.SetLabel(l.blockLabel)
		insts = append(insts, movInst)

	} else {
		// Store the value of the available register into the destination address
		strInst := arm.NewStr(fromR, arm.SP+", #"+destAddr)
		strInst.SetLabel(l.blockLabel)
		insts = append(insts, strInst)
	}

	return insts
}

// Build the stack table for the instruction.
func (l *Load) BuildStackTable(funcName string, stack *stack.Stack) {
	destinationReg := "r" + strconv.Itoa(l.targetRegisters[len(l.targetRegisters)-1])
	stack.AddEntry(funcName, destinationReg, strconv.Itoa(stack.GetFrame(funcName).GetLargestOffset()+8))
}
