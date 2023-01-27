package cmdline

import (
	"fmt"
	"os"
	"strings"
)

type Cmdline struct {
	Program string
	Args    []string
}

// Returns a new Cmdline struct with the program name and arguments
func newCmdline(program string, args []string) *Cmdline {
	return &Cmdline{
		Program: program,
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

	Args := os.Args[1:]

	cmdline := newCmdline(programFile, Args)

	return cmdline
}

// Check if the arguments are valid
func checkArgs() bool {
	return !(len(os.Args) < 2 || len(os.Args) > 3 || (len(os.Args) == 3 && os.Args[1] != "-lex"))
}

// Check if the file is a golite file
func isGoliteFile() bool {
	return strings.Split(os.Args[len(os.Args)-1], ".")[1] == "golite"
}

// usage of the program
func usage() {
	fmt.Println("Usage: go run golite/main.go [-lex] <input source file>\nArguments: <input source file> - path to the input source file\nFlags: -lex - Print the lexed tokens.")
}
