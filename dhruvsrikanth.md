# Compiler Report 

## Lexer

The lexer is the first stage of the compiler and is a part of the compiler frontend. The lexer takes in a character stream and outputs a token stream. This token stream is made up of lexed tokens which have been specified using regular expressions. Each token is associated with some amount of meta-information to allow for descriptive error messages. Once the regular expressions for each kind of token possible in the language have been specified, we may create a large combined NFA for all of the tokens. This NFA is then converted to a DFA which then enables us to find the *longest* token match for a given character stream. The way in which we have implemented the lexer utilizes a compiler frontend tool called [ANTLR](https://www.antlr.org/). Any errors generated at the lexing stage and recorded and propagated to the parser. Here is an overview of the different tokens possible for the language GoLite:

| Token | Regular Expression | Description |
| --- | --- | --- |
| `ASSIGN` | `=` | Assignment operator |
| `AND` | `&&` | Logical AND operator |
| `OR` | `\|\|` | Logical OR operator |
| `NOT` | `!` | Logical NOT operator |
| `EQ` | `==` | Equality operator |
| `NE` | `!=` | Inequality operator |
| `LT` | `<` | Less than operator |
| `LE` | `<=` | Less than or equal to operator |
| `GT` | `>` | Greater than operator |
| `GE` | `>=` | Greater than or equal to operator |
| `PLUS` | `+` | Addition operator |
| `MINUS` | `-` | Subtraction operator |
| `MULT` | `*` | Multiplication operator |
| `DIV` | `/` | Division operator |
| `LPAREN` | `(` | Left parenthesis |
| `RPAREN` | `)` | Right parenthesis |
| `LBRACE` | `{` | Left brace |
| `RBRACE` | `}` | Right brace |
| `SEMICOLON` | `;` | Semicolon |
| `COMMA` | `,` | Comma |
| `VAR` | `var` | Variable declaration |
| `TYPE` | `type` | Type declaration |
| `NEW` | `new` | New keyword |
| `DELETE` | `delete` | Delete keyword |
| `STRUCT` | `struct` | Struct keyword |
| `DOT` | `.` | Dot operator |
| `FUNC` | `func` | Function declaration |
| `RET` | `return` | Return keyword |
| `IF` | `if` | If keyword |
| `ELSE` | `else` | Else keyword |
| `FOR` | `for` | For keyword |
| `SCAN` | `scan` | Scan keyword |
| `PRINTF` | `printf` | Printf keyword |
| `INT_LIT` | `[1-9][0-9]* | [0]` | Integer literal |
| `STRING_LIT` | `'"'.*?'"'` | String literal |
| `BOOL_LIT` | `true | false` | Boolean literal |
| `NIL_LIT` | `nil` | Nil literal |
| `INT` | `int` | Integer type |
| `BOOL` | `bool` | Boolean type |
| `IDENT` | `[a-zA-Z][a-zA-Z0-9]*` | Identifier |
| `WS` | `[ \t\r\n]+` | Whitespace to skip |
| `COMMENT` | `'//'.*?'\n'` | Comment to skip |

## Parser

The parser is part of the latter half of a compiler frontend. The parser parses the token stream from the lexer to produce an AST (Abstract Syntax Tree). We use the [ANTLR](https://www.antlr.org/) tool to generate the parser, once more recording errors as they occur as well as accumulating any errors found at the lexing stage. The way in which we implement the parser follows a DFS-like tree visitor pattern. This is important as we will use it throughout the compiler to traverse our program data structures. It is used to create the AST, for type checking, while building the symbol table, for LLVM IR generation, and for ARM64 assembly code generation. In the parser, we build the AST by creating nodes for each of the tokens found in the token stream. The tree is created based on the grammar production rules provided below:

```antlr
Program = Types Declarations Functions 'eof'
Types = {TypeDeclaration}                   
TypeDeclaration = 'type' 'id' 'struct' '{' Fields '}' ';'                
Fields = Decl ';' {Decl ';'}                
Decl = 'id' Type                            
Type = 'int' | 'bool' | '*' 'id'            
Declarations = {Declaration}                
Declaration = 'var' Ids Type ';'            
Ids = 'id' {',' 'id'}                       
Functions = {Function}                      
Function = 'func' 'id' Parameters [ReturnType] '{' Declarations Statements '}'
Parameters = '(' [ Decl {',' Decl}] ')'     
ReturnType = type                           
Statements = {Statement}                    
Statement = Block | Assignment | Print | Delete | Conditional | Loop | Return | Read | Invocation 
Read =  'scan' LValue ';'                   
Block = '{' Statements '}'                  
Delete = 'delete' Expression ';'            
Assignment = LValue '=' Expression ';'      
Print = 'printf' '(' 'string' { ',' Expression} ')'  ';'                 
Conditional = 'if' '(' Expression ')' Block ['else' Block]               
Loop = 'for' '(' Expression ')' Block       
Return = 'return' [Expression] ';'          
Invocation = 'id' Arguments ';'             
Arguments = '(' [Expression {',' Expression}] ')'                        
LValue = 'id' {'.' id}                      
Expression = BoolTerm {'||' BoolTerm}       
BoolTerm = EqualTerm {'&&' EqualTerm}       
EqualTerm =  RelationTerm {('=='| '!=') RelationTerm}                    
RelationTerm = SimpleTerm {('>' | '<' | '<=' | '>=') SimpleTerm}         
SimpleTerm = Term {('+'| '-') Term}         
Term = UnaryTerm {('*'| '/') UnaryTerm}     
UnaryTerm = '!' SelectorTerm | '-' SelectorTerm | SelectorTerm           
SelectorTerm = Factor {'.' 'id'}            
Factor = '(' Expression ')' | 'id' [Arguments] | 'number' | 'new' 'id' | 'true' | 'false' | 'nil'
```

The tokens that represent terminal nodes or leaves in the production rules are adding to a sub-tree. When their parents are parsed, these nodes are added to the leaves of the parent nodes and the parent nodes are added to the AST. This happens in recursively in a bottom-up manner, building the tree from the leaves to the root (program node of the AST). This is bottom-up recursive tree traverse pattern is what we use thoughout the compiler. This can be accomplished by definining functions of the same name for each of the nodes in the data structure and then calling the appropriate function for the children of a node from within the parent node, then performing whatever work is need for the parent node involving the results of the children nodes. This is what makes the parser a *recursive descent parser*. In the lexing and parsing stages, we are able to catch syntax errors made in the program. As we will discuss later, the parser also consists of a semantic analysis stages where we catch semantic errors made by the programmer such as type mismatches between types used in an expression e.g. `1 + true` or using an undeclared variable.

## Static Semantics (Semantic Analysis) - Type Checking, Control Flow Checking and Symbol Tables

Static semantics deal with semantic analysis of the program. In order to perform efficient semantic analysis, we require a symbol table; a table containing the various global variable declarations, struct type declarations, function declarations (including the local variable declarations, parameters and return type). This includes meta information associated with these declarations such as type, scope and componenets (in the case of structs). Following a similar visitor pattern described above, we build up the symbol table by traversing the AST recursively down from the root to the leaves. Following this, we perform type checking on the program. This is done by using the same visitor pattern on the program node of the AST. This will call a function down the program paths from the root of the AST to the leaves. How we accomplish this is by calling a function `GetType()`. This function follows a similar pattern of first down to the leaves of each program path in the AST, and then propagating the inferred type of the component up the recursive path followed. We employ an clever method for type checking, wherein if there is a type mismatch at one level of the tree, we record this error, however, instead of propagating the incorrect type up the tree, we return the expected type, thereby eliminating repeated errors for the same type mismatch. At any given level of the tree, while performing type checking, we call `GetType()` on the respective expressions that we need to compare, then perform the comparison and record any errors that may occur. 

Another part of semantic analysis is associated with checking the control flow. **Every** function must have a valid control flow. This means that if a function's return type is `void`, it is not required for any `return` statement to be a part of the function. However, in the case that the return type of a function is not `void`, then every control path through the function must have a valid `return` statement. This is checked by traversing the AST in a similar visitor pattern described above and checking for the presence of a `return` statement at the end of each control path. We define a control flow path as the path a function may take starting from its local scope all the way down the call tree for each statement executed. We recursively traverse down this path and return `true` if we observe a `return` at some stage, otherwise `false`. Upon returning to the local scope and reaching the end of the function definition, if a `return` statement is not found, then we record an error. We do this for **every** control flow path found within a function.

Apart from these important aspects of semantic analysis, we also perform checks on the scope of variables used to see if they reside at a local scope, global scope or if they have not been declared, in which case an error is recorded.

As this falls under parsing, at the end of the parsing stage, we are returned an AST. We then build the symbol tables, perform semantic analysis and other forms of checking the validity of the program. We then proceed to the next stage of the compiler, if the program is valid. If node, we display the errors and exit the compiler. 

## Intermediate Representation (IR) - LLVM

We finally transition from the frontend of the compiler to the backend. This is largely made up of code generation, translations between intermediate representations and optimizations on the IRs. [LLVM](https://llvm.org/) is a compiler infrastructure that provides a set of tools and libraries for building compilers. It is a low-level virtual machine that provides a common intermediate representation (IR) for the compiler to work with. We utilized since static assignment i.e. each variable or register is assigned to at most once. We follwed a similar approach to AST construction for the creation of our LLVM IR. Traversing the AST in the same visitor pattern described previously, we build up the LLVM IR but translating AST nodes in to their IR representation. Each LLVM instruction was represented by a class or struct. Similar to the AST, we follow a recursive approach in building up the LLVM IR. During this process, we also convert the types listed in the symbol table to their LLVM IR counterparts, e.g. int -> i64. We store the entire LLVM IR in a similar structure to that of the AST, wherein the root node of the LLVM IR is the program node. Connected to this node, we have the global declarations, function declarations, runtime declarations (for C runtime functions such as `malloc` and `free`). Each function declaration node is made up of a control flow graph (CFG). The CFG enables us to dictate control flow through the function, which is especially useful for the `if` (`else`) and `for` statements. What was subtle and challenging about this stage was to effectively utilize pointers such that we are able to update the fields of a control flow block even if we don't know which block corresponds to it in a slice of control flow blocks as long as we have access to the pointer. Since a control flow block may eventually be split into several blocks (due to nested `if` or `for` statements), we need to evaluate the CFGs for the body of each statement before updating the predecessors, successors and branch labels for the current block. This was done with ease using pointers, however, I anticipate that this may be a challenge had we not been using pointers. Another interesting note is that a common reason for segmentation faults at this stage is due to incorrect branching. Check out this [this](llvm_ir_examples/) folder for some examples of LLVM IR generated by the compiler. This corresponds to the benchmarks found [here](benchmarks/simple/).

## Optimization

Optimizations deal with making the IR more efficient. The type of optimization is dependent of the goal of the optimization. It is possible to reduce the compile time which may be useful in real-time applications. It is also possible to reduce the run time which may be useful in HPC applications. There is tradeoff between runtime and compile time and therefore, optimizations must be chosen accordingly. I did **not** implement any optimizations, however, if I did I would have implemented the following optimization - dead code elimination. In this optimization, I would have focussed on the following aspects:

1. Removing code that is never executed.
2. Removing code that is redundant.

I would accomplish the first by traversing the CFG and removing any basic blocks that are not reachable from the entry block. The entry block here represents the first basic block in a function's CFG. This can be done by traversing the CFG and evaluating the condition for every conditional branch. If the condition can be fully evaluated at compile time i.e. it does not use any variables, then we can determine whether a branch will be taken or not. During the traverse, I can perform this exercise an mark all blocks as reachable or not. Following which, I can traverse the CFG to remove any unreachable blocks. Upon block removal, I can also remove any references i.e. by changing conditional branches to this block to unconditional branches to the branch other than the unreachable one in the conditional branch statement. We are able to obtain these references using the predecessors and successors of each block.

Similarly, I would traverse each block in a function's CFG to determine if there are any statements or declarations that are unused. For example, if a variable is declared but unsused throughout a function, then it can be eliminated. During the traverse of the CFGs, I would mark each LLVM instruction as used or unused. If an instruction is unused, then I would remove it along with any subsequent instructions involving the destination register (if any) of the LLVM instruction. We can do this because of the single static assignment methodology used in LLVM. Finally, I would do another traverse of the CFGs to remove any blocks that contain either no instructions or only a single unconditional branch instruction. The latter kind of block is prevalent in the CFGs produced after converting the AST to LLVM IR due to the canonical form of different types of control flows. The only updates to be made would be to record the destination label of the unconditional branch and update that block's predecessor blocks to use this destination label instead of the label of the block to be removed. In doing so we are able to eliminate any redundant blocks and dead code associated with the program.

