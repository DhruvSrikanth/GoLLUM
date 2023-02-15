package cmdline

import (
	"fmt"
	"os"
	"strings"
)

type Cmdline struct {
	Program string
	LexFlag bool
	ASTFlag bool
	Args    []string
}

// Returns a new Cmdline struct with the program name and arguments
func newCmdline(program string, lexFlag, astFlag bool, args []string) *Cmdline {
	return &Cmdline{
		Program: program,
		LexFlag: lexFlag,
		ASTFlag: astFlag,
		Args:    args,
	}
}

// ReadCmdline reads the kernel command line and returns a Cmdline struct.
func ReadCmdline() *Cmdline {
	if !checkArgs() {
		usage()
		return nil
	} else if !isGoliteFile() {
		fmt.Println("Input source file is not a golite file!")
		return nil
	}

	programFile := os.Args[0]
	var lexFlag bool
	var astFlag bool
	var args []string
	if len(os.Args) == 2 {
		lexFlag = false
		args = os.Args[1:]
	} else if len(os.Args) == 3 {
		if os.Args[1] == "-lex" {
			lexFlag = true
		} else if os.Args[1] == "-ast" {
			astFlag = true
		} else {
			usage()
			return nil
		}
		args = os.Args[2:]
	}

	cmdline := newCmdline(programFile, lexFlag, astFlag, args)

	return cmdline
}

// Check if the arguments are valid
func checkArgs() bool {
	if len(os.Args) < 2 {
		return false
	} else if len(os.Args) > 3 {
		return false
	} else if len(os.Args) == 3 {
		if os.Args[1] != "-lex" && os.Args[1] != "-ast" {
			return false
		}
	}
	return true
}

// Check if the file is a golite file
func isGoliteFile() bool {
	fileName := os.Args[len(os.Args)-1]
	return strings.HasSuffix(fileName, ".golite")
}

// usage of the program
func usage() {
	fmt.Println("Usage: go run golite/main.go [-lex || -ast] <input source file>\nArguments: <input source file> - path to the input source file\nFlags: \n-lex - Print the lexed tokens.\n-ast - Print the abstract syntax tree.")
}
