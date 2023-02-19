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

# Run symbol table tests
test_symbol_table:
	@make compiler
	@cd golite && \
	go test -v -run TestSymbolTable && \
	cd ..

# Run control flow tests
test_control_flow:
	@make compiler
	@cd golite && \
	go test -v -run TestControlFlow && \
	cd ..

# Run all parser tests
test_parser:
	@make compiler
	@make test_symbol_table
	@make test_type_checker
	@make test_control_flow

# Run all tests
test_compiler:
	@make test_lexer
	@make test_parser