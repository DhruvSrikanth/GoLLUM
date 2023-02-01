# Generate grammars
golite_compiler:
	@sh build.sh

compiler:
	@make --always-make golite_compiler

# Run simple correct tests
test_lexer:
	@make compiler
	@cd golite && \
	go test -v -run TestLexer && \
	cd ..

# Run simple incorrect tests
test_parser:
	@make compiler
	@cd golite && \
	go test -v -run TestParser && \
	cd ..

# Run all tests
test_compiler:
	@make test_lexer
	@make test_parser