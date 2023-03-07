package llvm

import (
	"bytes"
	"golite/arm"
	"golite/stack"
	"strconv"
	"strings"
)

// Represents accessing a struct field
type GetElementPtr struct {
	sourceRegister  string
	ty              string
	index           int
	blockLabel      string
	sourceRegisters []int
	targetRegisters []int
}

// NewGetElementPtr returns a new getelementptr instruction
func NewGetElementPtr(sourceRegister string, ty string, index int) *GetElementPtr {
	srcR := make([]int, 0)
	if strings.Contains(sourceRegister, "%r") {
		sourceRInt, _ := strconv.Atoi(sourceRegister[2:])
		srcR = append(srcR, sourceRInt)
	}

	tgtR := make([]int, 0)
	nextR, _ := strconv.Atoi(GetNextRegister()[2:])
	tgtR = append(tgtR, nextR)

	return &GetElementPtr{sourceRegister, ty, index, "", srcR, tgtR}
}

// String representation of the getelementptr instruction
func (g *GetElementPtr) String() string {
	var out bytes.Buffer

	// Format is - %result = getelementptr ty, ty* %source, i32 0, i32 index
	// Common for all
	out.WriteString("%r" + strconv.Itoa(g.targetRegisters[len(g.targetRegisters)-1]))
	out.WriteString(" = getelementptr ")
	// struct type
	if strings.Contains(g.ty, "struct.") {
		out.WriteString("%")
	}
	out.WriteString(g.ty)
	out.WriteString(", ")
	// struct type
	if strings.Contains(g.ty, "struct.") {
		out.WriteString("%")
	}
	out.WriteString(g.ty)
	out.WriteString("* ")
	out.WriteString(g.sourceRegister)
	out.WriteString(", i32 0, i32 ")
	out.WriteString(strconv.Itoa(g.index))

	return out.String()
}

// Get the registers targeted by the instruction.
func (g *GetElementPtr) GetTargets() []int {
	return g.targetRegisters
}

// Get the source registers of the instruction.
func (g *GetElementPtr) GetSources() []int {
	return g.sourceRegisters
}

// Get the immediate value of the instruction.
func (g *GetElementPtr) GetImmediate() int {
	return 0
}

// Get the label that marks this instruction in code.
func (g *GetElementPtr) GetLabel() string {
	return g.blockLabel
}

// Set the label that marks this instruction in code.
func (g *GetElementPtr) SetLabel(newLabel string) {
	g.blockLabel = newLabel
}

// Convert from LLVM IR to ARM assembly.
func (g *GetElementPtr) ToARM(fnName string, stack *stack.Stack) []arm.Instruction {
	insts := make([]arm.Instruction, 0)
	// We know that the source register is always a %r register in LLVM IR
	// First load the value of the source register (which is a pointer)
	// Then add the index * 8 to the value (which is a pointer)
	// This gives the address of the field we want to access
	// Now store the address i.e. the value of the source register + index * 8 in the target register
	stackFrame := stack.GetFrame(fnName)
	nextAvailableRegNum := stackFrame.GetNextRegister()
	nextAvailableReg := "x" + strconv.Itoa(nextAvailableRegNum)

	var srcR, destR string
	srcAddr := stackFrame.GetLocation(g.sourceRegister[1:])
	if strings.Contains(srcAddr, "x") {
		// The source register is in a register
		srcR = srcAddr
	} else {
		// The source register an address on the stack
		// Load the value of the source register into a register
		loadInst := arm.NewLdr(nextAvailableReg, arm.SP+", #"+srcAddr)
		loadInst.SetLabel(g.blockLabel)
		insts = append(insts, loadInst)

		srcR = nextAvailableReg
	}

	// Add the index * 8 to the value of the source register to get the address of the field we want to access
	offset := g.index * 8
	addInst := arm.NewAdd(srcR, srcR, "#"+strconv.Itoa(offset))
	addInst.SetLabel(g.blockLabel)
	insts = append(insts, addInst)

	// Get the destination register
	llvmDestReg := "r" + strconv.Itoa(g.targetRegisters[len(g.targetRegisters)-1])
	destAddr := stackFrame.GetLocation(llvmDestReg)
	if strings.Contains(destAddr, "x") {
		// The destination register is in a register
		destR = destAddr
		// Do a mov into the destination register
		movInst := arm.NewMov(destR, srcR)
		movInst.SetLabel(g.blockLabel)
		insts = append(insts, movInst)
	} else {
		// The destination register is on the stack
		// Store the value of the source register into the destination register
		strInst := arm.NewStr(srcR, arm.SP+", #"+destAddr)
		strInst.SetLabel(g.blockLabel)
		insts = append(insts, strInst)
	}
	return insts
}

// Build the stack table for the function.
func (g *GetElementPtr) BuildStackTable(funcName string, stack *stack.Stack) {
	// We need to indicate that the llvmDestReg is an address to the value and not the value itself
	// In a subequent LDR or STR instruction, we will have to use load it into a register, then use [reg] to indicate that it needs to be stored or loaded from the address
	destinationReg := "r" + strconv.Itoa(g.targetRegisters[len(g.targetRegisters)-1])
	stack.AddEntry(funcName, destinationReg, strconv.Itoa(stack.GetFrame(funcName).GetLargestOffset()+8), "pointer")
}
