# Define compilers used to test the project
LLVM_COMPILER="/opt/homebrew/opt/llvm/bin/llc"
ASSEMBLER="clang"
GO_COMPILER="go"

# Architecture to test
ARCH="x86_64-linux-gnu"
ARCH="arm64-apple-darwin22.2.0"

# Number of examples to test outputs on
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
	$(GO_COMPILER) test -v -run TestLexer && \
	cd ..

# Run type checker tests
test_type_checker:
	@make compiler
	@cd golite && \
	$(GO_COMPILER) test -v -run TestTypeChecker && \
	cd ..

# Run control flow tests
test_control_flow:
	@make compiler
	@cd golite && \
	$(GO_COMPILER) test -v -run TestControlFlow && \
	cd ..

# Run the ast tests
test_ast:
	@make compiler
	@cd golite && \
	$(GO_COMPILER) test -v -run TestAST && \
	cd ..

# Run all parser tests
test_parser:
	@make test_ast
	@make test_type_checker
	@make test_control_flow

# Run all compiler front end tests
test_frontend:
	@make test_lexer
	@make test_parser

# Run all compiler tests
test_compiler:
	@make test_frontend

# GENERATE
# Generate llvm for all examples
llvm_examples:
	@make compiler
	@for i in {1..$(N_EXAMPLES)} ; do \
        $(GO_COMPILER) run golite/main.go -llvm=$(ARCH) benchmarks/simple/example$$i.golite ; \
    done

# Generate assmebly for all examples (using llc)
assembly_examples_llvm:
	@make llvm_examples
	@for i in {1..$(N_EXAMPLES)} ; do \
		$(LLVM_COMPILER) example$$i.ll -o example$$i.s ; \
	done

# Generate assembly for all examples (using arm translation)
assembly_examples_arm:
	@for i in {1..$(N_EXAMPLES)} ; do \
		$(GO_COMPILER) run golite/main.go -s benchmarks/simple/example$$i.golite ; \
	done

# Generate object files for all examples (using clang)
llvm_examples_executable:
	@make assembly_examples_llvm
	@for i in {1..$(N_EXAMPLES)} ; do \
		$(ASSEMBLER) example$$i.s -o example$$i.out;\
	done

# Generate object files for all examples (using arm translation)
arm_examples_executable:
	@for i in {1..$(N_EXAMPLES)} ; do \
		$(GO_COMPILER) run golite/main.go -s -o example$$i.out benchmarks/simple/example$$i.golite ; \
	done