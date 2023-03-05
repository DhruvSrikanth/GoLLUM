package cmdline

import (
	"flag"
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
	ARMFlag         bool
	targetMachine   string
	Args            []string
}

// Returns a new Cmdline struct with the program name and arguments
func newCmdline(program string, lexFlag, astFlag, symbolTableFlag, LLVMShowFlag, ARMShowFlag, ARMFlag bool, targetMachine string, args []string) *Cmdline {
	return &Cmdline{
		Program:         program,
		LexFlag:         lexFlag,
		ASTFlag:         astFlag,
		SymbolTableFlag: symbolTableFlag,
		LLVMShowFlag:    LLVMShowFlag,
		ARMShowFlag:     ARMShowFlag,
		ARMFlag:         ARMFlag,
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
	if !isGoliteFile() {
		fmt.Println("Input source file is not a golite file!")
		return nil
	}

	programFile := os.Args[0]
	// targetMachine := "arm64-apple-darwin22.2.0"

	lexFlag := flag.Bool("lex", false, "Print the lexed tokens.")
	astFlag := flag.Bool("ast", false, "Print the abstract syntax tree.")
	symbolTableFlag := flag.Bool("sym", false, "Print the symbol table.")
	LLVMShowFlag := flag.Bool("llvmshow", false, "Print the LLVM IR representation.")
	ARMShowFlag := flag.Bool("arm64show", false, "Print the ARM64 assembly representation.")
	LLVMTargetTriple := flag.String("llvm", "x86_64-linux-gnu", "Generate the LLVM IR representation and output the LLVM IR to a file.")
	ARMFlag := flag.Bool("arm64", false, "Generate the ARM64 assembly representation and output the ARM64 assembly to a file.")

	flag.Parse()
	if strings.Contains(*LLVMTargetTriple, ".golite") {
		*LLVMTargetTriple = "x86_64-linux-gnu"
	}

	var args []string
	if len(os.Args) == 2 {
		args = os.Args[1:]
	} else if len(os.Args) == 3 {
		args = os.Args[2:]
	}

	cmdline := newCmdline(programFile, *lexFlag, *astFlag, *symbolTableFlag, *LLVMShowFlag, *ARMShowFlag, *ARMFlag, *LLVMTargetTriple, args)

	return cmdline
}

// Check if the file is a golite file
func isGoliteFile() bool {
	fileName := os.Args[len(os.Args)-1]
	return strings.HasSuffix(fileName, ".golite")
}

// usage of the program
func Usage() {
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
