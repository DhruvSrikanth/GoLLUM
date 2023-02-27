package llvm

import (
	"bytes"
	"fmt"
	"strings"
)

// This represents the target information to be included in the LLVM file.
type TargetInformation struct {
	fileName     string
	targetTriple string
}

// This creates a new target information object.
func NewTargetInformation(filePath, targetTriple string) *TargetInformation {
	filePathSplit := strings.Split(filePath, "/")
	file := filePathSplit[len(filePathSplit)-1]
	fileName := strings.Split(file, ".")[0]
	return &TargetInformation{fileName, targetTriple}
}

// String representation of the target information.
func (ti *TargetInformation) String() string {
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("source_filename = \"%s\"\n", ti.fileName))
	out.WriteString(fmt.Sprintf("target triple = \"%s\"\n\n", ti.targetTriple))
	return out.String()
}

// Get the file name.
func (ti *TargetInformation) GetFileName() string {
	return ti.fileName
}

// Get the target triple.
func (ti *TargetInformation) GetTargetTriple() string {
	return ti.targetTriple
}
