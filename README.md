
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
