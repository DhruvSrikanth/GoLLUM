package utils

import "os"

// Create the folder if it does not exist
func CreateFolder(folderPath string) {
	// Check if the folder exists
	if _, err := os.Stat(folderPath); err != nil {
		// Create the folder
		os.Mkdir(folderPath, os.ModePerm)
	}
}
