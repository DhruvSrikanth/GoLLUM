package cmdline

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Cmdline struct {
	Program          string
	LexFlag          bool
	ASTFlag          bool
	SymbolTableFlag  bool
	LLVMShowFlag     bool
	ARMShowFlag      bool
	ARMFlag          bool
	AssemblyFileName string
	targetMachine    string
	HelpFlag         bool
	InputProgram     string
}

// Returns a new Cmdline struct with the program name and arguments
func newCmdline(program string, lexFlag, astFlag, symbolTableFlag, LLVMShowFlag, ARMShowFlag, ARMFlag bool, assemblyFileName, targetMachine string, helpFlag bool, inputProgram string) *Cmdline {
	return &Cmdline{
		Program:          program,
		LexFlag:          lexFlag,
		ASTFlag:          astFlag,
		SymbolTableFlag:  symbolTableFlag,
		LLVMShowFlag:     LLVMShowFlag,
		ARMShowFlag:      ARMShowFlag,
		ARMFlag:          ARMFlag,
		AssemblyFileName: assemblyFileName,
		targetMachine:    targetMachine,
		HelpFlag:         helpFlag,
		InputProgram:     inputProgram,
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

	lexFlag := flag.Bool("lex", false, "Print the lexed tokens.")
	astFlag := flag.Bool("ast", false, "Print the abstract syntax tree.")
	symbolTableFlag := flag.Bool("sym", false, "Print the symbol table.")
	LLVMShowFlag := flag.Bool("llvmshow", false, "Print the LLVM IR representation.")
	ARMShowFlag := flag.Bool("arm64show", false, "Print the ARM64 assembly representation.")
	LLVMTargetTriple := flag.String("llvm", "x86_64-linux-gnu", "Generate the LLVM IR representation and output the LLVM IR to a file.")
	ARMFlag := flag.Bool("S", false, "Generate the ARM64 assembly representation and output the ARM64 assembly to a file.")
	AssemblyFileName := flag.String("o", "a.out", "Output assembly file name.")
	HelpFlag := flag.Bool("h", false, "Print the usage of the program.")

	flag.Parse()
	if strings.Contains(*LLVMTargetTriple, ".golite") {
		*LLVMTargetTriple = "x86_64-linux-gnu"
	}
	if strings.Contains(*AssemblyFileName, ".golite") {
		*AssemblyFileName = "a.out"
	}

	inputProgram := os.Args[len(os.Args)-1]

	cmdline := newCmdline(programFile, *lexFlag, *astFlag, *symbolTableFlag, *LLVMShowFlag, *ARMShowFlag, *ARMFlag, *AssemblyFileName, *LLVMTargetTriple, *HelpFlag, inputProgram)

	return cmdline
}

// Check if the file is a golite file
func isGoliteFile() bool {
	fileName := os.Args[len(os.Args)-1]
	return strings.HasSuffix(fileName, ".golite")
}

// usage of the program
func Usage() {
	usage := "Usage: go run golite/main.go [-lex || -ast || -sym || -llvmshow || -llvm=[<target triple>] || -arm64show || -s || -o] <input source file>\n"
	usage += "Arguments: <input source file> - path to the input source file\n"
	usage += "Flags: \n"
	usage += "\t-lex - Print the lexed tokens.\n"
	usage += "\t-ast - Print the abstract syntax tree.\n"
	usage += "\t-sym - Print the symbol table.\n"
	usage += "\t-llvmshow - Print the LLVM IR representation.\n"
	usage += "\t-arm64show - Print the ARM64 assembly representation.\n"
	usage += "\t-llvm=[<target triple>] - Generate the LLVM IR representation and output the LLVM IR to a file. Default target triple - x86_64-linux-gnu\n"
	usage += "\t-s -  Generate the ARM64 assembly representation and output the ARM64 assembly to a file.\n"
	usage += "\t-o - Output assembly file name. Default file name - a.out\n"
	usage += "\t-h - Print the usage of the program.\n"
	fmt.Println(usage)
}
