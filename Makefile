# Define compilers used to test the project
LLVM_COMPILER="/opt/homebrew/opt/llvm/bin/llc"
C_COMPILER="clang"

N_EXAMPLES=6

# BUILD
# Generate grammars
golite_compiler:
	@sh build.sh

compiler:
	@make --always-make golite_compiler

# TEST
# Run lexer tests
test_lexer:
	@make compiler
	@cd golite && \
	go test -v -run TestLexer && \
	cd ..

# Run type checker tests
test_type_checker:
	@make compiler
	@cd golite && \
	go test -v -run TestTypeChecker && \
	cd ..

# Run control flow tests
test_control_flow:
	@make compiler
	@cd golite && \
	go test -v -run TestControlFlow && \
	cd ..

# Run the ast tests
test_ast:
	@make compiler
	@cd golite && \
	go test -v -run TestAST && \
	cd ..

# Run all parser tests
test_parser:
	@make compiler
	@make test_ast
	@make test_type_checker
	@make test_control_flow

# Run all tests
test_compiler:
	@make test_lexer
	@make test_parser


# GENERATE
# Generate llvm for all examples
llvm_examples:
	@make compiler
	@for i in {1..$(N_EXAMPLES)} ; do \
        go run golite/main.go benchmarks/simple/example$$i.golite ; \
    done

# Generate assmebly for all examples (using llc)
assembly_examples_llvm:
	@make llvm_examples
	@for i in {1..$(N_EXAMPLES)} ; do \
		$(LLVM_COMPILER) IR/example$$i.ll -o assembly/example$$i.s ; \
	done
# Generate assembly for all examples (using arm translation)
assembly_examples_arm:
	@make llvm_examples
	@for i in {1..$(N_EXAMPLES)} ; do \
		go run golite/main.go -arm64 benchmarks/simple/example$$i.golite ; \
	done


# Generate object files for all examples (using clang)
llvm_examples_executable:
	@make assembly_examples_llvm
	@for i in {1..$(N_EXAMPLES)} ; do \
		$(C_COMPILER) assembly/example$$i.s -o example$$i.out;\
	done

# Generate object files for all examples (using arm translation)
arm_examples_executable:
	@make assembly_examples_arm
	@for i in {1..$(N_EXAMPLES)} ; do \
		$(C_COMPILER) assembly/example$$i.s -o example$$i.out;\
	done