package llvm

import (
	"bytes"
	"golite/arm"
	"golite/stack"
	"golite/utils"
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

	stackFrame := stack.GetFrame(fnName)
	availableRegNum := stackFrame.GetNextRegister()
	availableReg := "x" + strconv.Itoa(availableRegNum)

	var leftR, rightR string
	// Left register could be a literal, global variable, address or register
	// literal value => move into a register
	// global variable => load from address into a register
	// an address => load from address into a register
	// a register => use the register
	leftReg := b.leftRegister
	if strings.Contains(leftReg, "@") {
		// This is a global variable
		adrpInst := arm.NewAdrp(availableReg, leftReg[1:])
		adrpInst.SetLabel(b.blockLabel)
		insts = append(insts, adrpInst)

		addInst := arm.NewAdd(availableReg, availableReg, ":lo12:"+leftReg[1:])
		addInst.SetLabel(b.blockLabel)
		insts = append(insts, addInst)

		ldrInst := arm.NewLdr(availableReg, availableReg)
		ldrInst.SetLabel(b.blockLabel)
		insts = append(insts, ldrInst)

		leftR = availableReg
		availableRegNum += 1
	} else if strings.Contains(leftReg, "%") {
		leftAddr := stackFrame.GetLocation(leftReg[1:])
		if strings.Contains(leftReg, "x") {
			// This is a register
			leftR = leftReg[1:]
		} else {
			// This is an address
			ldrInst := arm.NewLdr(availableReg, arm.SP+", #"+leftAddr)
			ldrInst.SetLabel(b.blockLabel)
			insts = append(insts, ldrInst)

			leftR = availableReg
			availableRegNum += 1
		}
	} else {
		// This is a literal value
		// If its a negative value, we move 0 into a register, then subtract the value
		if strings.Contains(leftReg, "-") {
			movInst := arm.NewMov(availableReg, "#0")
			movInst.SetLabel(b.blockLabel)
			insts = append(insts, movInst)

			num, _ := strconv.Atoi(leftReg[1:])
			availableRegNum += 1
			nextReg := "x" + strconv.Itoa(availableRegNum)
			hexNum := utils.IntToHex(int64(num))
			hexSplit := utils.SplitHex(hexNum)
			// Move the upper 16 bits into the register and shift left by 48 bits
			movzInst := arm.NewMovz(nextReg, "#"+hexSplit[0]+", lsl #48")
			movzInst.SetLabel(b.blockLabel)
			insts = append(insts, movzInst)

			// Move the next 16 bits into the register and shift left by 32 bits
			movkInst := arm.NewMovk(nextReg, "#"+hexSplit[1]+", lsl #32")
			movkInst.SetLabel(b.blockLabel)
			insts = append(insts, movkInst)

			// Move the next 16 bits into the register and shift left by 16 bits
			movkInst = arm.NewMovk(nextReg, "#"+hexSplit[2]+", lsl #16")
			movkInst.SetLabel(b.blockLabel)
			insts = append(insts, movkInst)

			// Move the last 16 bits into the register
			movkInst = arm.NewMovk(nextReg, "#"+hexSplit[3]+", lsl #0")
			movkInst.SetLabel(b.blockLabel)
			insts = append(insts, movkInst)

			subInst := arm.NewSub(availableReg, availableReg, nextReg)
			subInst.SetLabel(b.blockLabel)
			insts = append(insts, subInst)

			leftR = availableReg
			availableRegNum += 1
		} else {
			// 4095 is the max value that we will move at a time
			if num, _ := strconv.Atoi(b.leftRegister); num > 4095 {
				hexNum := utils.IntToHex(int64(num))
				hexSplit := utils.SplitHex(hexNum)
				// Move the upper 16 bits into the register and shift left by 48 bits
				movzInst := arm.NewMovz(availableReg, "#"+hexSplit[0]+", lsl #48")
				movzInst.SetLabel(b.blockLabel)
				insts = append(insts, movzInst)

				// Move the next 16 bits into the register and shift left by 32 bits
				movkInst := arm.NewMovk(availableReg, "#"+hexSplit[1]+", lsl #32")
				movkInst.SetLabel(b.blockLabel)
				insts = append(insts, movkInst)

				// Move the next 16 bits into the register and shift left by 16 bits
				movkInst = arm.NewMovk(availableReg, "#"+hexSplit[2]+", lsl #16")
				movkInst.SetLabel(b.blockLabel)
				insts = append(insts, movkInst)

				// Move the last 16 bits into the register
				movkInst = arm.NewMovk(availableReg, "#"+hexSplit[3]+", lsl #0")
				movkInst.SetLabel(b.blockLabel)
				insts = append(insts, movkInst)
			} else {
				movInst := arm.NewMov(availableReg, "#"+b.leftRegister)
				movInst.SetLabel(b.blockLabel)
				insts = append(insts, movInst)
			}
			leftR = availableReg
			availableRegNum += 1
		}

	}

	// Right register could be a literal, global variable, address or register
	// literal value => move into a register
	// global variable => load from address into a register
	// an address => load from address into a register
	// a register => use the register
	availableReg = "x" + strconv.Itoa(availableRegNum)
	rightReg := b.rightRegister
	if strings.Contains(rightReg, "@") {
		// This is a global variable
		adrpInst := arm.NewAdrp(availableReg, rightReg[1:])
		adrpInst.SetLabel(b.blockLabel)
		insts = append(insts, adrpInst)

		addInst := arm.NewAdd(availableReg, availableReg, ":lo12:"+rightReg[1:])
		addInst.SetLabel(b.blockLabel)
		insts = append(insts, addInst)

		ldrInst := arm.NewLdr(availableReg, availableReg)
		ldrInst.SetLabel(b.blockLabel)
		insts = append(insts, ldrInst)

		rightR = availableReg
		availableRegNum += 1
	} else if strings.Contains(rightReg, "%") {
		rightAddr := stackFrame.GetLocation(rightReg[1:])
		if strings.Contains(rightReg, "x") {
			// This is a register
			rightR = rightReg[1:]
		} else {
			// This is an address
			ldrInst := arm.NewLdr(availableReg, arm.SP+", #"+rightAddr)
			ldrInst.SetLabel(b.blockLabel)
			insts = append(insts, ldrInst)

			rightR = availableReg
			availableRegNum += 1
		}
	} else {
		// This is a literal value
		if strings.Contains(rightReg, "-") {
			movInst := arm.NewMov(availableReg, "#0")
			movInst.SetLabel(b.blockLabel)
			insts = append(insts, movInst)

			num, _ := strconv.Atoi(rightReg[1:])
			// 4095 is the max value that we will move at a time
			availableRegNum += 1
			nextReg := "x" + strconv.Itoa(availableRegNum)
			hexNum := utils.IntToHex(int64(num))
			hexSplit := utils.SplitHex(hexNum)
			// Move the upper 16 bits into the register and shift left by 48 bits
			movzInst := arm.NewMovz(nextReg, "#"+hexSplit[0]+", lsl #48")
			movzInst.SetLabel(b.blockLabel)
			insts = append(insts, movzInst)

			// Move the next 16 bits into the register and shift left by 32 bits
			movkInst := arm.NewMovk(nextReg, "#"+hexSplit[1]+", lsl #32")
			movkInst.SetLabel(b.blockLabel)
			insts = append(insts, movkInst)

			// Move the next 16 bits into the register and shift left by 16 bits
			movkInst = arm.NewMovk(nextReg, "#"+hexSplit[2]+", lsl #16")
			movkInst.SetLabel(b.blockLabel)
			insts = append(insts, movkInst)

			// Move the last 16 bits into the register
			movkInst = arm.NewMovk(nextReg, "#"+hexSplit[3]+", lsl #0")
			movkInst.SetLabel(b.blockLabel)
			insts = append(insts, movkInst)

			subInst := arm.NewSub(availableReg, availableReg, nextReg)
			subInst.SetLabel(b.blockLabel)
			insts = append(insts, subInst)

			rightR = availableReg
			availableRegNum += 1
		} else {
			// 4095 is the max value that we will move at a time
			if num, _ := strconv.Atoi(b.rightRegister); num > 4095 {
				hexNum := utils.IntToHex(int64(num))
				hexSplit := utils.SplitHex(hexNum)
				// Move the upper 16 bits into the register and shift left by 48 bits
				movzInst := arm.NewMovz(availableReg, "#"+hexSplit[0]+", lsl #48")
				movzInst.SetLabel(b.blockLabel)
				insts = append(insts, movzInst)

				// Move the next 16 bits into the register and shift left by 32 bits
				movkInst := arm.NewMovk(availableReg, "#"+hexSplit[1]+", lsl #32")
				movkInst.SetLabel(b.blockLabel)
				insts = append(insts, movkInst)

				// Move the next 16 bits into the register and shift left by 16 bits
				movkInst = arm.NewMovk(availableReg, "#"+hexSplit[2]+", lsl #16")
				movkInst.SetLabel(b.blockLabel)
				insts = append(insts, movkInst)

				// Move the last 16 bits into the register
				movkInst = arm.NewMovk(availableReg, "#"+hexSplit[3]+", lsl #0")
				movkInst.SetLabel(b.blockLabel)
				insts = append(insts, movkInst)
			} else {
				movInst := arm.NewMov(availableReg, "#"+b.rightRegister)
				movInst.SetLabel(b.blockLabel)
				insts = append(insts, movInst)
			}

			rightR = availableReg
			availableRegNum += 1
		}
	}

	// operator can be
	// add, sub, mul, sdiv
	// and, or, eor
	// eq, neq, sgt, sge, slt, sle

	op := b.op
	isComparison := arm.IsComparisonOp(op)
	armOp := arm.LLVMOpToARMOp(op)
	availableReg = "x" + strconv.Itoa(availableRegNum)
	availableRegNum += 1
	if isComparison {
		// Do a comparison and set the flags (CMP)
		cmpInst := arm.NewCmp(leftR, rightR)
		cmpInst.SetLabel(b.blockLabel)
		insts = append(insts, cmpInst)

		// Store the result of the flag in the destintation register (CSET)
		csetInst := arm.NewCset(availableReg, armOp)
		csetInst.SetLabel(b.blockLabel)
		insts = append(insts, csetInst)
	} else {
		// Perform the normal operation
		binOpInst := arm.NewBinOp(availableReg, armOp, leftR, rightR)
		binOpInst.SetLabel(b.blockLabel)
		insts = append(insts, binOpInst)
	}

	// Destination register holds the result of the above operation
	var destR string
	destAddr := stackFrame.GetLocation("r" + strconv.Itoa(b.targetRegisters[len(b.targetRegisters)-1]))
	if strings.Contains(destAddr, "x") {
		// This is a register so we can move the result directly into it
		destR = destAddr

		movInst := arm.NewMov(destR, availableReg)
		movInst.SetLabel(b.blockLabel)
		insts = append(insts, movInst)

	} else {
		// This is an address so do a store
		strInst := arm.NewStr(availableReg, arm.SP+", #"+destAddr)
		strInst.SetLabel(b.blockLabel)
		insts = append(insts, strInst)
	}

	return insts
}

// Function to build the stack table for the BinOp.
func (b *BinOp) BuildStackTable(funcName string, stack *stack.Stack) {
	destinationReg := "r" + strconv.Itoa(b.targetRegisters[len(b.targetRegisters)-1])
	stack.AddEntry(funcName, destinationReg, strconv.Itoa(stack.GetFrame(funcName).GetLargestOffset()+8), "value")
}
