package cmdline

import (
	"os"
)

type Cmdline struct {
	Program string
	Args    []string
}

// Returns a new Cmdline struct with the program name and arguments
func NewCmdline(program string, args []string) *Cmdline {
	return &Cmdline{
		Program: program,
		Args:    args,
	}
}

// ReadCmdline reads the kernel command line and returns a Cmdline struct.
func ReadCmdline() *Cmdline {
	programFile := os.Args[0]

	Args := os.Args[1:]

	cmdline := newCmdline(programFile, Args)

	return cmdline
}
