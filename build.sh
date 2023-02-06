#!/bin/bash
go mod tidy

# Generate the lexer and parser

cd grammars 
source generate.sh 
cd ..