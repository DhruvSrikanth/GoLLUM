package ast

import (
	"fmt"
	"golite/token"
)

// Semantic Analysis Error
type SemanticAnalysisError struct {
	Message string       // Message to display
	Type    string       // Type of error (redeclaration, invalid type)
	token   *token.Token // Token where the error occurred
}

// New Semantic Analysis Error
func NewSemanticAnalysisError(message, errorType string, tok *token.Token) *SemanticAnalysisError {
	return &SemanticAnalysisError{message, errorType, tok}
}

// String representation of the Semantic Analysis Error
func (e *SemanticAnalysisError) String() string {
	return fmt.Sprintf("semantic error at line (%d), col (%d): %s (%s)", e.token.GetLine(), e.token.GetColumn(), e.Message, e.Type)
}
