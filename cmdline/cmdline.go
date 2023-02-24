package cmdline

import (
	"fmt"
	"os"
	"strings"
)

type Cmdline struct {
	Program         string
	LexFlag         bool
	ASTFlag         bool
	SymbolTableFlag bool
	LLVMFlag        bool
	Args            []string
}

// Returns a new Cmdline struct with the program name and arguments
func newCmdline(program string, lexFlag, astFlag, symbolTableFlag, llvmFlag bool, args []string) *Cmdline {
	return &Cmdline{
		Program:         program,
		LexFlag:         lexFlag,
		ASTFlag:         astFlag,
		SymbolTableFlag: symbolTableFlag,
		LLVMFlag:        llvmFlag,
		Args:            args,
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
	var symbolTableFlag bool
	var llvmFlag bool
	var args []string
	if len(os.Args) == 2 {
		lexFlag = false
		args = os.Args[1:]
	} else if len(os.Args) == 3 {
		if os.Args[1] == "-lex" {
			lexFlag = true
		} else if os.Args[1] == "-ast" {
			astFlag = true
		} else if os.Args[1] == "-sym" {
			symbolTableFlag = true
		} else if os.Args[1] == "-llvm" {
			llvmFlag = true
		} else {
			usage()
			return nil
		}
		args = os.Args[2:]
	}

	cmdline := newCmdline(programFile, lexFlag, astFlag, symbolTableFlag, llvmFlag, args)

	return cmdline
}

// Check if the arguments are valid
func checkArgs() bool {
	if len(os.Args) < 2 {
		return false
	} else if len(os.Args) > 3 {
		return false
	} else if len(os.Args) == 3 {
		if os.Args[1] != "-lex" && os.Args[1] != "-ast" && os.Args[1] != "-sym" && os.Args[1] != "-llvm" {
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
	usage := "Usage: go run golite/main.go [-lex || -ast || -sym || -llvm] <input source file>\n"
	usage += "Arguments: <input source file> - path to the input source file\n"
	usage += "Flags: \n-lex - Print the lexed tokens.\n-ast - Print the abstract syntax tree.\n-sym - Print the symbol table.\n-llvm - Print the LLVM IR representation.\n"
	fmt.Println(usage)
}
