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
	LLVMShowFlag    bool
	ARMShowFlag     bool
	targetMachine   string
	Args            []string
}

// Returns a new Cmdline struct with the program name and arguments
func newCmdline(program string, lexFlag, astFlag, symbolTableFlag, LLVMShowFlag, ARMShowFlag bool, targetMachine string, args []string) *Cmdline {
	return &Cmdline{
		Program:         program,
		LexFlag:         lexFlag,
		ASTFlag:         astFlag,
		SymbolTableFlag: symbolTableFlag,
		LLVMShowFlag:    LLVMShowFlag,
		ARMShowFlag:     ARMShowFlag,
		targetMachine:   targetMachine,
		Args:            args,
	}
}

// Get the target triple
func (c *Cmdline) GetTargetTriple() string {
	return c.targetMachine
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
	// targetMachine := "arm64-apple-darwin22.2.0"
	targetMachine := "x86_64-linux-gnu"

	var lexFlag bool
	var astFlag bool
	var symbolTableFlag bool
	var llvmShowFlag bool
	var ARMShowFlag bool
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
		} else if os.Args[1] == "-llvmshow" {
			llvmShowFlag = true
		} else if os.Args[1] == "-arm64show" {
			ARMShowFlag = true
		} else {
			usage()
			return nil
		}
		args = os.Args[2:]
	}

	cmdline := newCmdline(programFile, lexFlag, astFlag, symbolTableFlag, llvmShowFlag, ARMShowFlag, targetMachine, args)

	return cmdline
}

// Check if the arguments are valid
func checkArgs() bool {
	if len(os.Args) < 2 {
		return false
	} else if len(os.Args) > 3 {
		return false
	} else if len(os.Args) == 3 {
		if os.Args[1] != "-lex" && os.Args[1] != "-ast" && os.Args[1] != "-sym" && os.Args[1] != "-llvmshow" && os.Args[1] != "-arm64show" && os.Args[1] != "-llvm" && os.Args[1] != "-arm64" {
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
	usage := "Usage: go run golite/main.go [-lex || -ast || -sym || -llvmshow || llvm=[<target triple>] || arm64show || arm64] <input source file>\n"
	usage += "Arguments: <input source file> - path to the input source file\n"
	usage += "Flags: \n-lex - Print the lexed tokens.\n"
	usage += "-ast - Print the abstract syntax tree.\n"
	usage += "-sym - Print the symbol table.\n"
	usage += "-llvmshow - Print the LLVM IR representation.\n"
	usage += "-arm64show - Print the ARM64 assembly representation.\n"
	usage += "-llvm=[<target triple>] - Generate the LLVM IR representation and output the LLVM IR to a file. Default target triple - x86_64-linux-gnu\n"
	usage += "-arm64 -  Generate the ARM64 assembly representation and output the ARM64 assembly to a file.\n"
	fmt.Println(usage)
}
