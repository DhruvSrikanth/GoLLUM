
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
go run golite/main.go benchmarks/simple/simple.golite 
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
go run golite/main.go benchmarks/bad/bad.golite 
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
