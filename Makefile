# Generate grammars
golite:
	source build.sh

# Run simple correct tests
test_correct:
	make golite
	go run golite/main.go benchmarks/simple/simple.golite 
	go run golite/main.go -lex benchmarks/simple/simple.golite 

# Run simple incorrect tests
test_incorrect:
	make golite
	go run golite/main.go benchmarks/bad/bad.golite 
	go run golite/main.go -lex benchmarks/bad/bad.golite 

# Run all tests
test:
	make golite
	go run golite/main.go benchmarks/simple/simple.golite 
	go run golite/main.go -lex benchmarks/simple/simple.golite 
	go run golite/main.go benchmarks/bad/bad.golite 
	go run golite/main.go -lex benchmarks/bad/bad.golite 