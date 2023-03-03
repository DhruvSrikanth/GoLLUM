package llvm

import "fmt"

// Global counter variable for labels.
var labelCounter int

// Get the next label (gives the block to be added).
func GetNextLabel() string {
	nextLabel := fmt.Sprintf("L%d", labelCounter)
	labelCounter++
	return nextLabel
}

// Get the current label (gives the block that will be added next).
func GetCurrentLabel() string {
	return fmt.Sprintf("L%d", labelCounter)
}

// Get the previous label (gives the last block already added).
func GetPreviousLabel() string {
	return fmt.Sprintf("L%d", labelCounter-1)
}

// Global counter variable for registers.
var registerCounter int

// Get the next register.
func GetNextRegister() string {
	nextReg := fmt.Sprintf("%%r%d", registerCounter)
	registerCounter++
	return nextReg
}

// Get the current register.
func GetCurrentRegister() string {
	return fmt.Sprintf("%%r%d", registerCounter)
}

// Get the previous register.
func GetPreviousRegister() string {
	return fmt.Sprintf("%%r%d", registerCounter-1)
}
