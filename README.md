
<p align="center">
<a>
    <img src="https://readme-typing-svg.demolab.com?font=Georgia&size=28&duration=3500&pause=2000&multiline=true&width=1000&height=80&lines=GoLLUM - Go + Lite + Language + Understanding + Machine" alt="Typing SVG" />
</a>
<br/>

# GoLLUM

## Build Golite compiler

```bash
source build.sh
```

or 

```bash
make compiler
```

## Run Golite compiler

Example of code that will work:

```Go
func main() {
    
    var a int;
    
    a = 3 + 4 + 5;
    
    printf("%d", a);

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

This runs both the `lexer` and `parser`. If you want to only run the `lexer` and print the tokens (along with errors found along the way) lexed, use the `-lex` flag. For example - 

```bash
go run golite/main.go -lex benchmarks/simple/simple1.golite 
```

## Run tests using Makefile

Build `golite`:
```shell
make compiler
```

Test `lexer`:
```shell
make test_lexer
```

Test `parser`:
```shell
make test_parser
```


Test all completed portions of the compiler:
```shell
make test_compiler
```

Running the tests present in `golite/main_test.go` tests the compiler in both the `lexer` and `parser` phases. Ideally this should return something like - 
```bash
=== RUN   TestLexer
=== RUN   TestLexer/T=1
=== RUN   TestLexer/T=2
=== RUN   TestLexer/T=3
=== RUN   TestLexer/T=4
=== RUN   TestLexer/T=5
=== RUN   TestLexer/T=6
=== RUN   TestLexer/T=7
=== RUN   TestLexer/T=8
=== RUN   TestLexer/T=9
=== RUN   TestLexer/T=10
--- PASS: TestLexer (3.81s)
    --- PASS: TestLexer/T=1 (0.46s)
    --- PASS: TestLexer/T=2 (0.24s)
    --- PASS: TestLexer/T=3 (0.27s)
    --- PASS: TestLexer/T=4 (0.54s)
    --- PASS: TestLexer/T=5 (0.24s)
    --- PASS: TestLexer/T=6 (0.27s)
    --- PASS: TestLexer/T=7 (0.54s)
    --- PASS: TestLexer/T=8 (0.23s)
    --- PASS: TestLexer/T=9 (0.51s)
    --- PASS: TestLexer/T=10 (0.51s)
PASS
ok      golite/golite   4.325s
=== RUN   TestParser
=== RUN   TestParser/T=1
=== RUN   TestParser/T=2
=== RUN   TestParser/T=3
=== RUN   TestParser/T=4
=== RUN   TestParser/T=5
=== RUN   TestParser/T=6
=== RUN   TestParser/T=7
=== RUN   TestParser/T=8
=== RUN   TestParser/T=9
--- PASS: TestParser (1.96s)
    --- PASS: TestParser/T=1 (0.21s)
    --- PASS: TestParser/T=2 (0.31s)
    --- PASS: TestParser/T=3 (0.23s)
    --- PASS: TestParser/T=4 (0.19s)
    --- PASS: TestParser/T=5 (0.19s)
    --- PASS: TestParser/T=6 (0.24s)
    --- PASS: TestParser/T=7 (0.19s)
    --- PASS: TestParser/T=8 (0.20s)
    --- PASS: TestParser/T=9 (0.20s)
PASS
ok      golite/golite   2.064s
```
