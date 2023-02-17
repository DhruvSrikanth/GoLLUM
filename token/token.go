package token

type Token struct {
	Line   int // The line where this token was found
	Column int // The column where this token was found
}

func NewToken(line, column int) *Token {
	return &Token{line, column}
}

// Get line number
func (t *Token) GetLine() int {
	return t.Line
}

// Get column number
func (t *Token) GetColumn() int {
	return t.Column
}
