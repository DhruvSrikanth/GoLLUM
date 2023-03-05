package utils

import "os"

// Write the representation to the output file
func WriteRepr(outputPath, llvmRepr string) {
	// Output path is within the current working directory
	// Create the output file
	outFile, err := os.Create(outputPath)
	if err != nil {
		return
	}
	defer outFile.Close()

	// Write the representation to the output file
	outFile.WriteString(llvmRepr)
}

// Function to remove the representation file created
func RemoveRepr(outputPath string) {
	// Check if the file exists
	if _, err := os.Stat(outputPath); err == nil {
		// Remove the file
		os.Remove(outputPath)
	}
}
