package ast

import "fmt"

// Semantic Analysis Error
type SemanticAnalysisError struct {
	Message string // Message to display
	Type    string // Type of error (redeclaration, invalid type)
}

// New Semantic Analysis Error
func NewSemanticAnalysisError(message string, errorType string) *SemanticAnalysisError {
	return &SemanticAnalysisError{message, errorType}
}

// String representation of the Semantic Analysis Error
func (e *SemanticAnalysisError) String() string {
	return fmt.Sprintf("semantic error: %s (%s)", e.Message, e.Type)
}
