
<p align="center">
<a>
    <img src="https://readme-typing-svg.demolab.com?font=Georgia&size=28&duration=3500&pause=2000&multiline=true&width=1000&height=80&lines=GoLLUM - Go + Lite + Language + Understanding + Machine" alt="Typing SVG" />
</a>
<br/>

# GoLLUM

![build](https://img.shields.io/badge/build-passing-brightgreen)
![coverage](https://img.shields.io/badge/coverage-100%25-brightgreen)
![Go](https://img.shields.io/badge/Go-1.18-blue)
![LLVM](https://img.shields.io/badge/LLVM-15.0.7__1-red)
![Clang](https://img.shields.io/badge/Clang-15.0.7-green)

## Requirements

- [Go](https://golang.org/dl/) 1.18 or higher
- [LLVM](https://llvm.org/) 15.0.7_1 or higher (only required for running LLVM tests or building executables with LLVM backend)
- [Clang](https://clang.llvm.org/)

## Build Golite Compiler

```bash
make compiler
```

## Run Golite Compiler

```bash
go run golite/main.go benchmarks/simple/example2.golite 
```

## Golite Compiler Flags

| Flag | Description | Default | Stages | Example |
| --- | --- | --- | --- | --- |
| `h` | Print the program usage and flags. | False | None | `go run golite/main.go -h benchmarks/simple/example1.golite` |
| `lex` | Print the lexed tokens. | False | Lexer | `go run golite/main.go -lex benchmarks/simple/example1.golite` |
| `ast` | Print the parser produced by the parser. | False |  Lexer, Parser | `go run golite/main.go -ast benchmarks/simple/example1.golite` |
| `sym` | Print the symbol table. | False | Lexer, Parser | `go run golite/main.go -sym benchmarks/simple/example1.golite` |
| `llvmshow` | Print the LLVM IR. | False | Lexer, Parser, IR Generator | `go run golite/main.go -llvmshow benchmarks/simple/example1.golite` |
| `llvm=[target triple]` | Generate LLVM IR using target triple. | x86_64-linux-gnu | Lexer, Parser, IR Generator | `go run golite/main.go -llvm=arm64-apple-darwin22.2.0 benchmarks/simple/example1.golite` |
| `arm64show` | Print the ARM64 assembly. | False | Lexer, Parser, IR Generator, ARM64 Generator | `go run golite/main.go -arm64show benchmarks/simple/example1.golite` |
| `S` | Generate ARM64 assembly. | True | Lexer, Parser, IR Generator, ARM64 Generator | `go run golite/main.go -arm64 benchmarks/simple/example1.golite` |
| `o [output file]` | Output file name. | `a.out` | Lexer, Parser, IR Generator, ARM64 Generator, Assembler | `go run golite/main.go -o output benchmarks/simple/example1.golite` |

Check out the `benchmarks` folder for more examples.

Important to note:

- Executables produced by the compiler along with any intermediate files requested through flags are generated in the root directory of the project.

## Test Golite Compiler

| Command | Description |
| --- | --- |
| `make test_lexer` | Runs lexer tests. |
| `make test_ast` | Runs the ast construction tests. |
| `make test_type_checking` | Runs type checker tests. |
| `make test_control_flow` | Runs control flow tests. |
| `make test_parser` | Runs all parser tests. |
| `make test_frontend` | Runs frontend tests. |
| `make test_llvm` | Runs LLVM IR generator tests. |
| `make test_arm` | Runs ARM64 generator tests. |
| `make backend` | Runs backend tests. |
| `make test_compiler` | Runs all compiler tests. |

Running the tests present in `golite/main_test.go` will run all the tests in the `benchmarks` folder.

Important to note:
- To run any LLVM IR tests, you will need to have LLVM installed on your machine.

## Examples of Golite Programs

Examples of Golite programs can be found in the `benchmarks/simple` folder.

Example of code that will work:

```Go
type Node struct { 

    value int; 
    left  *Node;
    right *Node;  
}; 

var  root *Node; 

func compare(cur int,neuw int) int {
    if (cur < neuw) {
		return 1;
	} else {  
		if (cur > neuw) {
		   return -1;
		} else {
		   return 0;
		}
	}
}

func addNode(numAdd int, curr *Node) { 
	var compVal int;
    var newNode *Node;

	if (curr == nil){
		newNode = new Node;
		newNode.value = numAdd;
		root = newNode;
	} else {
		compVal = compare(curr.value, numAdd);
		
		if (compVal == -1) {
			if (curr.left == nil) {
				newNode = new Node;
				newNode.value = numAdd;
				curr.left = newNode;
			} else {
				addNode(numAdd, curr.left);
			}
		} else {
			if (compVal == 1) {
				if (curr.right == nil) {
					newNode = new Node;
					newNode.value = numAdd;
					curr.right = newNode;
				} else {
					addNode(numAdd, curr.right);
				}
			}
		}
	}
}

func printDepthTree(curr *Node)  {
    var temp int;
	if (curr != nil) {
		if (curr.left != nil) {
			printDepthTree(curr.left);
		}
        temp = curr.value;
		printf("%d",temp);

		if (curr.right !=nil) {
			printDepthTree(curr.right);
		}
	}
}

func deleteLeavesTree(curr *Node) {
	if (curr != nil) {
		if (curr.left != nil) {
			deleteLeavesTree(curr.left);
		}
		
		if (curr.right !=nil) {
			deleteLeavesTree(curr.right);
		}
		delete(curr);
	}
}

func main() {
	var input, temp int;

	root = nil;
	input = 0;

	scan input; 

	for (input!=0) {
		addNode(input, root);
		scan input; 
	}
	
	printDepthTree(root);	
	deleteLeavesTree(root);
}
```

Example of code that will not work:

```Go
func main() {
    var a bool;
    var x_ double;
    a = +;
    printf("%d", a);
}
```
