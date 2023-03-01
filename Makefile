LLVM_PATH="/opt/homebrew/opt/llvm/bin/llc"

# Generate grammars
golite_compiler:
	@sh build.sh

compiler:
	@make --always-make golite_compiler

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