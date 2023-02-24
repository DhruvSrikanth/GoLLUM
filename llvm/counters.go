package llvm

import "fmt"

// Global counter variable for labels.
var labelCounter int

// Get the next label.
func GetNextLabel() string {
	nextLabel := fmt.Sprintf("label-%d", labelCounter)
	labelCounter++
	return nextLabel
}

// Global counter variable for registers.
var registerCounter int

// Get the next register.
func GetNextRegister() string {
	nextReg := fmt.Sprintf("%%r%d", registerCounter)
	registerCounter++
	return nextReg
}
