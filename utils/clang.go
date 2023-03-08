package utils

import (
	"os/exec"
)

// Generate the executable file
func GenerateExecutable(assemblyPath, outputFileName string) {
	// Create the executable file using clang
	cmd := exec.Command("clang", "-o", outputFileName, assemblyPath)
	cmd.Run()
}
