
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
go run golite/main.go -lex benchmarks/simple/simple.golite 
```

Example of code that will not work:

```Go
func main() {
    
    var a bool;
    
    a = +;
    
    printf("%d", a);

}
```

```bash
go run golite/main.go -lex benchmarks/bad/bad.golite 
```

## Run tests using Makefile

Build `golite`:
```shell
make golite
```

Correct program tests:
```shell
make test_correct
```

Incorrect program tests:
```shell
make test_incorrect
```

All tests:
```shell
make test
```
