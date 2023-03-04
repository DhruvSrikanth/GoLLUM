package llvm

import (
	"bytes"
	"golite/arm"
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
func (s *Store) ToARM() []*arm.Instruction {
	return nil
}
