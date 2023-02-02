package main

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func stringSliceEquals(expected, actual []string) bool {
	if len(expected) != len(actual) {
		return false
	}
	for i, v := range expected {
		if v != actual[i] {
			return false
		}
	}
	return true
}

func formatOutput(out string) []string {
	s := strings.TrimSpace(out)
	actual := strings.Split(s, "\n")
	return actual

}

func TestLexer(t *testing.T) {
	lexFlag := "-lex"

	var lexTests = []struct {
		programPath string
		expected    []string
	}{
		// Good programs
		{"../benchmarks/simple/simple1.golite", []string{"func", "main", "(", ")", "{", "var", "a", "int", ";", "a", "=", "3", "+", "4", "+", "5", ";", "printf", "(", "\"%d\"", ",", "a", ")", ";", "}", "<EOF>"}},
		{"../benchmarks/simple/simple2.golite", []string{"func", "main", "(", ")", "{", "var", "a", "int", ";", "var", "b", "int", ";", "b", "=", "5", ";", "a", "=", "1", ";", "for", "(", "a", "<", "10", ")", "{", "a", "=", "a", "+", "1", ";", "if", "(", "a", "==", "b", ")", "{", "printf", "(", "\"a == b\"", ")", ";", "}", "if", "(", "a", "!=", "b", ")", "{", "printf", "(", "\"a != b\"", ")", ";", "}", "if", "(", "a", "<", "b", ")", "{", "printf", "(", "\"a < b\"", ")", ";", "}", "if", "(", "a", ">", "b", ")", "{", "printf", "(", "\"a > b\"", ")", ";", "}", "if", "(", "a", "<=", "b", ")", "{", "printf", "(", "\"a <= b\"", ")", ";", "}", "if", "(", "a", ">=", "b", ")", "{", "printf", "(", "\"a >= b\"", ")", ";", "}", "}", "}", "<EOF>"}},
		{"../benchmarks/simple/simple3.golite", []string{"func", "main", "(", ")", "{", "var", "ptr", "*", "a", ";", "ptr", "=", "new", "a", ";", "delete", "ptr", ";", "}", "<EOF>"}},
		{"../benchmarks/simple/simple4.golite", []string{"func", "main", "(", ")", "{", "var", "n", "int", ";", "var", "result", "int", ";", "n", "=", "scan", ";", "result", "=", "fib", "(", "n", ")", ";", "printf", "(", "\"fibonacci(%d)= %d\"", ",", "n", ",", "result", ")", ";", "}", "func", "fib", "(", "n", "int", ")", "int", "{", "var", "temp1", ",", "temp2", "int", ";", "var", "res1", ",", "res2", ",", "result", "int", ";", "temp1", "=", "n", "-", "1", ";", "temp2", "=", "n", "-", "2", ";", "if", "(", "n", "==", "0", ")", "{", "return", "0", ";", "}", "else", "{", "if", "(", "n", "==", "1", ")", "{", "return", "1", ";", "}", "else", "{", "res1", "=", "fib", "(", "temp1", ")", ";", "res2", "=", "fib", "(", "temp2", ")", ";", "result", "=", "res1", "+", "res2", ";", "return", "result", ";", "}", "}", "}", "<EOF>"}},
		{"../benchmarks/simple/simple5.golite", []string{"func", "main", "(", ")", "{", "type", "Point", "struct", "{", "x", "int", ";", "y", "int", ";", "}", ";", "var", "p1", "Point", ";", "var", "p2", "Point", ";", "var", "dx", ",", "dy", "int", ";", "var", "l2Squared", "int", ";", "p1", ".", "x", "=", "10", ";", "p1", ".", "y", "=", "20", ";", "p2", ".", "x", "=", "30", ";", "p2", ".", "y", "=", "40", ";", "dx", "=", "p1", ".", "x", "-", "p2", ".", "x", ";", "dy", "=", "p1", ".", "y", "-", "p2", ".", "y", ";", "l2Squared", "=", "dx", "*", "dx", "+", "dy", "*", "dy", ";", "printf", "(", "\"The squared distance between the two points is %d\"", ",", "l2Squared", ")", ";", "return", ";", "}", "<EOF>"}},

		// Bad programs
		{"../benchmarks/bad/bad1.golite", []string{"func", "main", "(", ")", "{", "var", "a", "bool", ";", "var", "x", "double", ";", "a", "=", "+", ";", "printf", "(", "\"%d\"", ",", "a", ")", ";", "}", "<EOF>", "lexer error(4:9): token recognition error at: '_'"}},
		{"../benchmarks/bad/bad2.golite", []string{"func", "main", "(", ")", "{", "var", "a", "int", ";", "var", "b", "int", ";", "var", "c", "int", ";", "printf", "(", "\"Enter a number: \"", ")", ";", "a", "=", "scan", ";", "printf", "(", "\"Enter another number: \"", ")", ";", "b", "=", "scan", ";", "c", "=", "(", "a", "+", "b", ")", ")", ";", "printf", "(", "\"The sum of the two numbers is: \"", ",", "c", ")", ";", "}", "<EOF>"}},
		{"../benchmarks/bad/bad3.golite", []string{"func", "main", "(", ")", "{", "var", "a", "bool", ";", "var", "b", "bool", ";", "var", "c", "bool", ";", "var", "d", "bool", ";", "var", "e", "bool", ";", "a", "=", "true", ";", "b", "=", "false", ";", "c", "=", "a", "&&", "b", ";", "d", "=", "a", "||", "b", ";", "e", "=", "!", "a", ";", "f", "=", "(", "a", "&&", "b", ")", "||", "(", "c", "and", "d", ")", "||", "e", ";", "printf", "(", "\"%b %b %b %b %b %b\"", ",", "a", ",", "b", ",", "c", ",", "d", ",", "e", ",", "f", ")", ";", "}", "<EOF>"}},
		{"../benchmarks/bad/bad4.golite", []string{"func", "main", "(", ")", "{", "var", "a", "int", ";", "var", "i", "int", ";", "i", "=", "0", ";", "a", "{", "i", "}", "=", "1", ";", "}", "<EOF>"}},
		{"../benchmarks/bad/bad5.golite", []string{"func", "main", "(", ")", "{", "var", "a", ",", "b", ",", "c", "int", ";", "return", "sum", "(", "a", ",", "b", ",", "c", ")", ";", "}", "func", "sum", "(", "a", ",", "b", ",", "c", "int", ")", "{", "return", "a", "+", "b", "+", "c", ";", "}", "<EOF>"}},
	}
	for num, test := range lexTests {
		testname := fmt.Sprintf("T=%v", num+1)
		t.Run(testname, func(t *testing.T) {
			var err error
			cmd := exec.Command("go", "run", "main.go", lexFlag, test.programPath)
			out, err := cmd.Output()
			actual := formatOutput(string(out))
			if err != nil || !stringSliceEquals(test.expected, actual) {
				t.Errorf("\nRan:%s\nExpected:%s\nGot:%s", cmd, test.expected, actual)
			}
		})
	}
}

func TestParser(t *testing.T) {
	var parseTests = []struct {
		programPath string
		expected    []string
	}{
		// Good programs
		{"../benchmarks/simple/simple1.golite", []string{""}},
		{"../benchmarks/simple/simple2.golite", []string{""}},
		{"../benchmarks/simple/simple3.golite", []string{""}},
		{"../benchmarks/simple/simple4.golite", []string{""}},
		{"../benchmarks/simple/simple5.golite", []string{""}},

		// Bad programs
		{"../benchmarks/bad/bad1.golite", []string{"lexer error(4:9): token recognition error at: '_'"}},
		{"../benchmarks/bad/bad2.golite", []string{"syntax error(10,15): extraneous input ')' expecting ';'"}},
		{"../benchmarks/bad/bad3.golite", []string{"syntax error(13,23): mismatched input 'and' expecting ')'"}},
		{"../benchmarks/bad/bad4.golite", []string{"syntax error(5,5): no viable alternative at input 'a{'", "syntax error(5,7): no viable alternative at input 'i}'", "syntax error(5,9): mismatched input '=' expecting '}'"}},
		{"../benchmarks/bad/bad5.golite", []string{"syntax error(6,10): mismatched input ',' expecting {'*', 'int', 'bool'}", "syntax error(6,13): mismatched input ',' expecting {'*', 'int', 'bool'}"}},
	}
	for num, test := range parseTests {
		testname := fmt.Sprintf("T=%v", num+1)
		t.Run(testname, func(t *testing.T) {
			var err error
			cmd := exec.Command("go", "run", "main.go", test.programPath)
			out, err := cmd.Output()
			actual := formatOutput(string(out))
			if err != nil || !stringSliceEquals(test.expected, actual) {
				t.Errorf("\nRan:%s\nExpected:%s\nGot:%s", cmd, test.expected, actual)
			}
		})
	}
}
