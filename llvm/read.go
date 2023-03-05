package llvm

import (
	"bytes"
	"golite/arm"
	"golite/stack"
	"strconv"
	"strings"
)

// Representation of a Read runtime call
type Read struct {
	sourceReg       string
	blockLabel      string
	sourceRegisters []int
	targetRegisters []int
}

// NewRead returns a new Read runtime call
func NewRead(sourceRegister string) *Read {
	srcR := make([]int, 0)
	if strings.Contains(sourceRegister, "%r") {
		rInt, _ := strconv.Atoi(sourceRegister[2:])
		srcR = append(srcR, rInt)
	}

	tgtR := make([]int, 0)

	return &Read{
		sourceReg:       sourceRegister,
		blockLabel:      "",
		sourceRegisters: srcR,
		targetRegisters: tgtR,
	}
}

// String representation of the Read call
func (r *Read) String() string {
	var out bytes.Buffer
	// Format is -  call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %register)
	// Common for all
	out.WriteString("call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* ")
	out.WriteString(r.sourceReg)
	out.WriteString(")")

	return out.String()
}

// Get the registers targeted by the instruction.
func (r *Read) GetTargets() []int {
	return r.targetRegisters
}

// Get the source registers of the instruction.
func (r *Read) GetSources() []int {
	return r.sourceRegisters
}

// Get the immediate value of the instruction.
func (r *Read) GetImmediate() int {
	return 0
}

// Get the label that marks this instruction in code.
func (r *Read) GetLabel() string {
	return r.blockLabel
}

// Set the label that marks this instruction in code.
func (r *Read) SetLabel(newLabel string) {
	r.blockLabel = newLabel
}

// Convert LLVM IR to ARM assembly.
func (r *Read) ToARM(fnName string, stack *stack.Stack) []arm.Instruction {
	insts := make([]arm.Instruction, 0)

	stackFrame := stack.GetFrame(fnName)

	// Compute the address of the string using the global label
	// and store it in the next available register
	availableRegNum := stackFrame.GetNextRegister()
	availableReg := "x" + strconv.Itoa(availableRegNum)
	availableRegNum += 1
	adrpInst := arm.NewAdrp(availableReg, ".READ")
	adrpInst.SetLabel(r.blockLabel)
	insts = append(insts, adrpInst)

	addInst := arm.NewAdd(availableReg, availableReg, ":lo12:.READ")
	addInst.SetLabel(r.blockLabel)
	insts = append(insts, addInst)

	// Move this into x0
	movInst := arm.NewMov("x0", availableReg)
	movInst.SetLabel(r.blockLabel)
	insts = append(insts, movInst)

	if strings.Contains(r.sourceReg, "@") {
		availableReg = "x" + strconv.Itoa(availableRegNum)
		src := r.sourceReg[1:]
		// Global variable that needs to be address found in a different way
		// Get the address
		adrpInst := arm.NewAdrp(availableReg, src)
		adrpInst.SetLabel(r.blockLabel)
		insts = append(insts, adrpInst)
		// Get the full address
		addInst := arm.NewAdd(availableReg, availableReg, ":lo12:"+src)
		addInst.SetLabel(r.blockLabel)
		insts = append(insts, addInst)
		// Move this into x1
		movInst := arm.NewMov("x1", availableReg)
		movInst.SetLabel(r.blockLabel)
		insts = append(insts, movInst)

	} else {
		src := r.sourceReg[1:]
		destAddr := stackFrame.GetLocation(src)
		if strings.Contains(destAddr, "x") {
			// Do a move into x1
			movInst := arm.NewMov("x1", destAddr)
			movInst.SetLabel(r.blockLabel)
			insts = append(insts, movInst)
		} else {
			// Move the address into a register
			addInst := arm.NewAdd("x1", arm.SP, "#"+destAddr)
			addInst.SetLabel(r.blockLabel)
			insts = append(insts, addInst)
		}

	}
	return insts
}

// Build the stack table for the instruction.
func (r *Read) BuildStackTable(fName string, stack *stack.Stack) {
}
