package llvm

import (
	"bytes"
	"fmt"
)

// This represents the target information to be included in the LLVM file.
type TargetInformation struct {
	fileName     string
	targetTriple string
}

// This creates a new target information object.
func NewTargetInformation(fileName, targetTriple string) *TargetInformation {
	return &TargetInformation{fileName, targetTriple}
}

// String representation of the target information.
func (ti *TargetInformation) String() string {
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("source_filename = \"%s\"", ti.fileName))
	out.WriteString(fmt.Sprintf("target triple = \"%s\"", ti.targetTriple))
	return out.String()
}
