#!/bin/bash


alias antlr4='java -Xmx500M -cp "./antlr-4.11.1-complete.jar:$CLASSPATH" org.antlr.v4.Tool'

#Please note this code produces Golang generated code you will need ot change the flag -Dlanguage to your specific language of choice. 
antlr4 -Dlanguage=Go -no-visitor -package lexer -o ../lexer GoliteLexer.g4
antlr4 -Dlanguage=Go -no-visitor -package parser -lib ../lexer -o ../parser GoliteParser.g4