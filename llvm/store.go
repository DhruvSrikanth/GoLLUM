package llvm

import (
	"bytes"
	"golite/arm"
	"golite/stack"
	"strconv"
	"strings"
)

// Represents the store LLVM IR instruction.
type Store struct {
	sourceRegister  string
	targetRegister  string
	ty              string
	blockLabel      string
	sourceRegisters []int
	targetRegisters []int
}

// Create a new store instruction.
func NewStore(sourceRegister, targetRegister, ty string) *Store {
	srcR := make([]int, 0)
	if !strings.Contains(sourceRegister, "@") && !strings.Contains(sourceRegister, "%r") {
		// This is a literal value
	} else {
		// This is a register
		if strings.Contains(sourceRegister, "%r") {
			srcReg, _ := strconv.Atoi(sourceRegister[2:])
			srcR = append(srcR, srcReg)
		}
	}

	tgtR := make([]int, 0)
	if strings.Contains(targetRegister, "%r") {
		tgtReg, _ := strconv.Atoi(targetRegister[2:])
		tgtR = append(tgtR, tgtReg)
	}
	return &Store{sourceRegister, targetRegister, ty, "", srcR, tgtR}
}

// String representation of the store instruction.
func (s *Store) String() string {
	var out bytes.Buffer
	// Format is - store ty value, ty* register
	// ty can be a struct type or i64
	// Common for all
	out.WriteString("store ")
	// struct type
	if strings.Contains(s.ty, "struct.") {
		out.WriteString("%")
	}
	// Common for all
	out.WriteString(s.ty)
	out.WriteString(" ")
	out.WriteString(s.sourceRegister)
	out.WriteString(", ")
	// struct type
	if strings.Contains(s.ty, "struct.") {
		out.WriteString("%")
	}
	// Common for all
	out.WriteString(s.ty)
	out.WriteString("* ")
	out.WriteString(s.targetRegister)

	return out.String()
}

// Get the registers targeted by the instruction.
func (s *Store) GetTargets() []int {
	return s.targetRegisters
}

// Get the source registers of the instruction.
func (s *Store) GetSources() []int {
	return s.sourceRegisters
}

// Get the label of the block containing the instruction.
func (s *Store) GetLabel() string {
	return s.blockLabel
}

// Get the immediate value of the instruction.
func (s *Store) GetImmediate() int {
	return 0
}

// Set the label of the block containing the instruction.
func (s *Store) SetLabel(label string) {
	s.blockLabel = label
}

// Convert from LLVM IR to ARM assembly.
func (s *Store) ToARM(fnName string, stack *stack.Stack) []arm.Instruction {
	insts := make([]arm.Instruction, 0)

	stackFrame := stack.GetFrame(fnName)
	availableRegNum := stackFrame.GetNextRegister()
	availableReg := "x" + strconv.Itoa(availableRegNum)

	var srcR, destR string
	// The source could be a literal value, a global variable, a register or an address
	// literal value => move into a register
	// global variable => load from address into a register
	// an address => load from address into a register
	// a register => use the register
	if strings.Contains(s.sourceRegister, "@") {
		// This is a global variable
		adrpInst := arm.NewAdrp(availableReg, s.sourceRegister[1:])
		adrpInst.SetLabel(s.blockLabel)
		insts = append(insts, adrpInst)

		addInst := arm.NewAdd(availableReg, availableReg, ":lo12:"+s.sourceRegister[1:])
		addInst.SetLabel(s.blockLabel)
		insts = append(insts, addInst)

		ldrInst := arm.NewLdr(availableReg, availableReg)
		ldrInst.SetLabel(s.blockLabel)
		insts = append(insts, ldrInst)

		srcR = availableReg
		availableRegNum += 1

	} else if strings.Contains(s.sourceRegister, "%") {
		// Check the stack frame for the register/address
		srcAddr := stackFrame.GetLocation(s.sourceRegister[1:])
		if strings.Contains(srcAddr, "x") {
			// Register
			srcR = srcAddr
		} else {
			// Address
			ldrInst := arm.NewLdr(availableReg, arm.SP+", #"+srcAddr)
			ldrInst.SetLabel(s.blockLabel)
			insts = append(insts, ldrInst)

			srcR = availableReg
			availableRegNum += 1
		}
	} else {
		// This is a literal value
		// If its a negative value, mov 0 into a register and then sub the value
		if strings.Contains(s.sourceRegister, "-") {
			movInst := arm.NewMov(availableReg, "#0")
			movInst.SetLabel(s.blockLabel)
			insts = append(insts, movInst)

			num, _ := strconv.Atoi(s.sourceRegister[1:])
			// 4095 is the max value that we will move at a time
			for num > 4095 {
				subInst := arm.NewSub(availableReg, availableReg, "#4095")
				subInst.SetLabel(s.blockLabel)
				insts = append(insts, subInst)

				num -= 4095
			}
			remaining := strconv.Itoa(num)
			subInst := arm.NewSub(availableReg, availableReg, "#"+remaining)
			subInst.SetLabel(s.blockLabel)
			insts = append(insts, subInst)
		} else {
			// 4095 is the max value that we will move at a time
			if num, _ := strconv.Atoi(s.sourceRegister); num > 4095 {
				movInst := arm.NewMov(availableReg, "#4095")
				movInst.SetLabel(s.blockLabel)
				insts = append(insts, movInst)
				remaining := num - 4095
				for remaining > 4095 {
					addInst := arm.NewAdd(availableReg, availableReg, "#4095")
					addInst.SetLabel(s.blockLabel)
					insts = append(insts, addInst)

					remaining -= 4095
				}
				remainingStr := strconv.Itoa(remaining)
				addInst := arm.NewAdd(availableReg, availableReg, "#"+remainingStr)
				addInst.SetLabel(s.blockLabel)
				insts = append(insts, addInst)
			} else {
				movInst := arm.NewMov(availableReg, "#"+s.sourceRegister)
				movInst.SetLabel(s.blockLabel)
				insts = append(insts, movInst)
			}
		}

		srcR = availableReg
		availableRegNum += 1
	}

	// The destination could be a global variable, a register, an address on the stack, or a pointer to an address
	// global variable => use the address with a store
	// register => do a move instead of a store
	// address => use the address with a store
	// pointer to an address => load the address of the pointer which is the value of the register and then do a store into that address i.e. [loaded reg]

	availableReg = "x" + strconv.Itoa(availableRegNum)
	if strings.Contains(s.targetRegister, "@") {
		// This is a global variable
		adrpInst := arm.NewAdrp(availableReg, s.targetRegister[1:])
		adrpInst.SetLabel(s.blockLabel)
		insts = append(insts, adrpInst)

		addInst := arm.NewAdd(availableReg, availableReg, ":lo12:"+s.targetRegister[1:])
		addInst.SetLabel(s.blockLabel)
		insts = append(insts, addInst)

		destR = availableReg

		strInst := arm.NewStr(srcR, destR)
		strInst.SetLabel(s.blockLabel)
		insts = append(insts, strInst)
	} else {
		destAddr := stackFrame.GetLocation(s.targetRegister[1:])
		destAddrTy := stackFrame.GetType(s.targetRegister[1:])
		if strings.Contains(destAddr, "x") {
			// Register so do a mov with the srcR
			movInst := arm.NewMov(destAddr, srcR)
			movInst.SetLabel(s.blockLabel)
			insts = append(insts, movInst)
			return insts
		} else {
			// Address so use the offset from sp as the address to store in
			if destAddrTy == "value" {
				destR = arm.SP + ", #" + destAddr
				strInst := arm.NewStr(srcR, destR)
				strInst.SetLabel(s.blockLabel)
				insts = append(insts, strInst)
			} else {
				// This is a pointer so we need to load the address of the pointer and then store the value at that address
				availableReg = "x" + strconv.Itoa(availableRegNum+1)
				ldrInst := arm.NewLdr(availableReg, arm.SP+", #"+destAddr)
				ldrInst.SetLabel(s.blockLabel)
				insts = append(insts, ldrInst)

				strInst := arm.NewStr(srcR, availableReg)
				strInst.SetLabel(s.blockLabel)
				insts = append(insts, strInst)
			}
		}
	}

	return insts
}

// Build the stack table for the instruction.
func (s *Store) BuildStackTable(fName string, stack *stack.Stack) {
}
