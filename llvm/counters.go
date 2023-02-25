package llvm

import "fmt"

// Global counter variable for labels.
var labelCounter int

// Get the next label.
func GetNextLabel() string {
	nextLabel := fmt.Sprintf("L%d", labelCounter)
	labelCounter++
	return nextLabel
}

// Get the current label.
func GetCurrentLabel() string {
	return fmt.Sprintf("L%d", labelCounter)
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
