
<p align="center">
<a>
    <img src="https://readme-typing-svg.demolab.com?font=Georgia&size=28&duration=3500&pause=2000&multiline=true&width=1000&height=80&lines=GoLLUM - Go + Lite + Language + Understanding + Machine" alt="Typing SVG" />
</a>
<br/>

# GoLLUM

## Build Golite Compiler

```bash
make compiler
```

## Run Golite Compiler

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
func addNode (numAdd int, curr *Node) { 
	var compVal int;
    var newNode *Node;

	if( curr == nil ){
		newNode = new Node;
		newNode.value = numAdd;
		root = newNode;
	} else {
		compVal = compare (curr.value, numAdd);
		
		if( compVal == -1) {
			if(curr.left == nil) {
				newNode = new Node;
				newNode.value = numAdd;
				curr.left = newNode;
			} else {
				addNode ( numAdd, curr.left);
			}
		} else {
			if(compVal == 1) {
				if(curr.right == nil) {
					newNode = new Node;
					newNode.value = numAdd;
					curr.right = newNode;
				} else {
					addNode ( numAdd, curr.right);
				}
			}
		}
	}
}
func printDepthTree (curr *Node)  {
    var temp int;
	if( curr != nil ) {
		if( curr.left != nil) {
			printDepthTree(curr.left);
		}
        temp = curr.value;
		printf("%d",temp);

		if(curr.right !=nil) {
			printDepthTree(curr.right);
		}
	}
}
func deleteLeavesTree (curr *Node) {
	if( curr != nil ) {
		if( curr.left != nil) {
			deleteLeavesTree(curr.left);
		}

		if(curr.right !=nil) {
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
		addNode (input, root);
		scan input; 
	}
	
	printDepthTree(root);	
	deleteLeavesTree(root);
}
```

Run - 

```bash
go run golite/main.go benchmarks/simple/simple1.golite 
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

```bash
go run golite/main.go benchmarks/bad/bad1.golite 
```

This runs both the `lexer` and `parser`. 

## Golite Compiler Flags

| Flag | Description | Example |
| --- | --- | --- |
| -lex | Print the lexed tokens (Runs lexer). | `go run golite/main.go -lex benchmarks/example/example1.golite` |
| -ast | Print the parser produced by the parser (Runs lexer and parser). | `go run golite/main.go -ast benchmarks/example/example1.golite` |
| -sym | Print the symbol table (Runs lexer and parser). | `go run golite/main.go -sym benchmarks/example/example1.golite` |

Check out the `benchmarks` folder for more examples.

## Test Golite Compiler

| Command | Description |
| --- | --- |
| `make test_lexer` | Runs lexer tests. |
| `make test_ast` | Runs the ast construction tests. |
| `make test_type_checking` | Runs type checker tests. |
| `make test_control_flow` | Runs control flow tests. |
| `make test_parser` | Runs all parser tests. |
| `make test_compiler` | Runs all compiler tests. |


Running the tests present in `golite/main_test.go` tests the compiler in both the `lexer` and `parser` phases.