
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
    
    a int;
    
    a = 3 + 4 + 5;
    
    printf("%d", a);

}
```

```bash
go run golite/main.go -lex benchmarks/bad/bad.golite 
```

# Open Questions:
1. Questions on Ed
2. Questions in the lexer script
3. Questions in the parser script
4. Questions about the print, string and "%d" in the printf function
5. Questions about build not working but generate working
6. Questions about the -lex flag and which part of the code it needs to be implemented in and how to get the parser to stop parsing and how to get all the tokens
7. Is error checking fine and some additional print statements such as the one at the end of main.go
8. Can i use a makefile instead of the build and generate files