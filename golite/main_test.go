package main

import (
	"fmt"
	"golite/utils"
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
		{"../benchmarks/lexer/good/good1.golite", []string{"func", "main", "(", ")", "{", "var", "a", "int", ";", "a", "=", "3", "+", "4", "+", "5", ";", "printf", "(", "\"%d\"", ",", "a", ")", ";", "}", "<EOF>"}},
		{"../benchmarks/lexer/good/good2.golite", []string{"func", "main", "(", ")", "{", "var", "a", "int", ";", "var", "b", "int", ";", "b", "=", "5", ";", "a", "=", "1", ";", "for", "(", "a", "<", "10", ")", "{", "a", "=", "a", "+", "1", ";", "if", "(", "a", "==", "b", ")", "{", "printf", "(", "\"a == b\"", ")", ";", "}", "if", "(", "a", "!=", "b", ")", "{", "printf", "(", "\"a != b\"", ")", ";", "}", "if", "(", "a", "<", "b", ")", "{", "printf", "(", "\"a < b\"", ")", ";", "}", "if", "(", "a", ">", "b", ")", "{", "printf", "(", "\"a > b\"", ")", ";", "}", "if", "(", "a", "<=", "b", ")", "{", "printf", "(", "\"a <= b\"", ")", ";", "}", "if", "(", "a", ">=", "b", ")", "{", "printf", "(", "\"a >= b\"", ")", ";", "}", "}", "}", "<EOF>"}},
		{"../benchmarks/lexer/good/good3.golite", []string{"func", "main", "(", ")", "{", "var", "ptr", "*", "a", ";", "ptr", "=", "new", "a", ";", "delete", "ptr", ";", "}", "<EOF>"}},
		{"../benchmarks/lexer/good/good4.golite", []string{"func", "main", "(", ")", "{", "var", "n", "int", ";", "var", "result", "int", ";", "n", "=", "scan", ";", "result", "=", "fib", "(", "n", ")", ";", "printf", "(", "\"fibonacci(%d)= %d\"", ",", "n", ",", "result", ")", ";", "}", "func", "fib", "(", "n", "int", ")", "int", "{", "var", "temp1", ",", "temp2", "int", ";", "var", "res1", ",", "res2", ",", "result", "int", ";", "temp1", "=", "n", "-", "1", ";", "temp2", "=", "n", "-", "2", ";", "if", "(", "n", "==", "0", ")", "{", "return", "0", ";", "}", "else", "{", "if", "(", "n", "==", "1", ")", "{", "return", "1", ";", "}", "else", "{", "res1", "=", "fib", "(", "temp1", ")", ";", "res2", "=", "fib", "(", "temp2", ")", ";", "result", "=", "res1", "+", "res2", ";", "return", "result", ";", "}", "}", "}", "<EOF>"}},
		{"../benchmarks/lexer/good/good5.golite", []string{"type", "Point", "struct", "{", "x", "int", ";", "y", "int", ";", "}", ";", "func", "main", "(", ")", "{", "var", "p1", "*", "Point", ";", "var", "p2", "*", "Point", ";", "var", "result", "int", ";", "p1", "=", "new", "Point", ";", "p2", "=", "new", "Point", ";", "p1", ".", "x", ".", "y", "=", "10", ";", "result", "=", "distance2D", "(", "p1", ",", "p2", ")", ";", "printf", "(", "\"The squared distance between the two points is %d\"", ",", "l2Squared", ")", ";", "delete", "p1", ";", "delete", "p2", ";", "return", ";", "}", "func", "distance2D", "(", "p1", "*", "Point", ",", "p2", "*", "Point", ")", "int", "{", "var", "dx", ",", "dy", "int", ";", "var", "l2Squared", "int", ";", "l2Squared", "=", "dx", "*", "dx", "+", "dy", "*", "dy", ";", "return", "l2Squared", ";", "}", "<EOF>"}},

		// Bad programs
		{"../benchmarks/lexer/bad/bad1.golite", []string{"func", "main", "(", ")", "{", "var", "a", "bool", ";", "var", "x", "double", ";", "a", "=", "+", ";", "printf", "(", "\"%d\"", ",", "a", ")", ";", "}", "<EOF>", "lexer error(4:9): token recognition error at: '_'"}},
		{"../benchmarks/lexer/bad/bad2.golite", []string{"func", "main", "(", ")", "{", "var", "a", "int", ";", "var", "b", "int", ";", "var", "c", "int", ";", "printf", "(", "\"Enter a number: \"", ")", ";", "a", "=", "scan", ";", "printf", "(", "\"Enter another number: \"", ")", ";", "b", "=", "scan", ";", "c", "=", "(", "a", "+", "b", ")", ")", ";", "printf", "(", "\"The sum of the two numbers is: \"", ",", "c", ")", ";", "}", "<EOF>"}},
		{"../benchmarks/lexer/bad/bad3.golite", []string{"func", "main", "(", ")", "{", "var", "a", "bool", ";", "var", "b", "bool", ";", "var", "c", "bool", ";", "var", "d", "bool", ";", "var", "e", "bool", ";", "a", "=", "true", ";", "b", "=", "false", ";", "c", "=", "a", "&&", "b", ";", "d", "=", "a", "||", "b", ";", "e", "=", "!", "a", ";", "f", "=", "(", "a", "&&", "b", ")", "||", "(", "c", "and", "d", ")", "||", "e", ";", "printf", "(", "\"%b %b %b %b %b %b\"", ",", "a", ",", "b", ",", "c", ",", "d", ",", "e", ",", "f", ")", ";", "}", "<EOF>"}},
		{"../benchmarks/lexer/bad/bad4.golite", []string{"func", "main", "(", ")", "{", "var", "a", "int", ";", "var", "i", "int", ";", "i", "=", "0", ";", "a", "{", "i", "}", "=", "1", ";", "}", "<EOF>"}},
		{"../benchmarks/lexer/bad/bad5.golite", []string{"func", "main", "(", ")", "{", "var", "a", ",", "b", ",", "c", "int", ";", "return", "sum", "(", "a", ",", "b", ",", "c", ")", ";", "}", "func", "sum", "(", "a", ",", "b", ",", "c", "int", ")", "{", "return", "a", "+", "b", "+", "c", ";", "}", "<EOF>"}},
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

func TestAST(t *testing.T) {
	var parseTests = []struct {
		programPath string
		expected    []string
	}{
		// Good programs
		{"../benchmarks/parser/ast/good/good1.golite", []string{"type nums struct {", "a int;", "b int;", "}", "func fib1(f1 int) int {", "if ((f1 < 2)) {", "return f1;", "} else {", "return (fib1((f1 - 1)) + fib1((f1 - 2)));", "}", "}", "func fib2(f2 int) int {", "var fir int;", "var sec int;", "var temp int;", "fir = 0;", "sec = 1;", "for ((f2 != 0)) {", "f2 = (f2 - 1);", "temp = (fir + sec);", "fir = sec;", "sec = temp;", "}", "return fir;", "}", "func main() void {", "var x *nums;", "var c int;", "var d int;", "var temp int;", "x = new nums;", "scan temp;", "x.a = temp;", "scan temp;", "x.b = temp;", "c = fib1(x.a);", "d = fib2(x.b);", "delete x;", "}"}},
		{"../benchmarks/parser/ast/good/good2.golite", []string{"func isqrt(a int) int {", "var square int;", "var delta int;", "square = 1;", "delta = 3;", "for ((square <= a)) {", "square = (square + delta);", "delta = (delta + 2);", "}", "return ((delta / 2) - 1);", "}", "func prime(a int) bool {", "var max int;", "var divisor int;", "var remainder int;", "if ((a < 2)) {", "return false;", "} else {", "max = isqrt(a);", "divisor = 2;", "for ((divisor <= max)) {", "remainder = (a - ((a / divisor) * divisor));", "if ((remainder == 0)) {", "return false;", "}", "divisor = (divisor + 1);", "}", "return true;", "}", "}", "func main() void {", "var limit int;", "var a int;", "scan limit;", "a = 0;", "for ((a <= limit)) {", "if (prime(a)) {", "printf(\"%d\", a);", "}", "a = (a + 1);", "}", "}"}},
	}
	for num, test := range parseTests {
		testname := fmt.Sprintf("T=%v", num+1)
		t.Run(testname, func(t *testing.T) {
			var err error
			cmd := exec.Command("go", "run", "main.go", "-ast", test.programPath)
			out, err := cmd.Output()
			actual := formatOutput(string(out))
			if err != nil || !stringSliceEquals(test.expected, actual) {
				t.Errorf("\nRan:%s\nExpected:%s\nGot:%s", cmd, test.expected, actual)
			}
		})
	}
}

func TestControlFlow(t *testing.T) {
	var parseTests = []struct {
		programPath string
		expected    []string
	}{
		// Good programs
		{"../benchmarks/parser/control flow/good/good1.golite", []string{""}},
		{"../benchmarks/parser/control flow/good/good2.golite", []string{""}},
		{"../benchmarks/parser/control flow/good/good3.golite", []string{""}},
		{"../benchmarks/parser/control flow/good/good4.golite", []string{""}},
		{"../benchmarks/parser/control flow/good/good5.golite", []string{""}},
		{"../benchmarks/parser/control flow/good/good6.golite", []string{""}},
		{"../benchmarks/parser/control flow/good/good7.golite", []string{""}},

		// Bad programs
		{"../benchmarks/parser/control flow/bad/bad1.golite", []string{"semantic error at line (-1), col (-1): main function not found (main function error)", "semantic error at line (7), col (0): control flow does not reach return statement (invalid control flow)", "Failed semantic analysis"}},
		{"../benchmarks/parser/control flow/bad/bad2.golite", []string{"semantic error at line (-1), col (-1): main function not found (main function error)", "semantic error at line (2), col (0): control flow does not reach return statement (invalid control flow)", "Failed semantic analysis"}},
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

	// Remove the IR created from the tests
	utils.RemoveFolder("IR")
	// Remove the assembly folder created from the tests
	utils.RemoveFolder("assembly")
}

func TestTypeChecker(t *testing.T) {
	var parseTests = []struct {
		programPath string
		expected    []string
	}{
		// Good programs
		{"../benchmarks/parser/type checking/good/good1.golite", []string{""}},
		{"../benchmarks/parser/type checking/good/good2.golite", []string{""}},
		{"../benchmarks/parser/type checking/good/good3.golite", []string{""}},
		{"../benchmarks/parser/type checking/good/good4.golite", []string{""}},
		{"../benchmarks/parser/type checking/good/good5.golite", []string{""}},
		{"../benchmarks/parser/type checking/good/good6.golite", []string{""}},
		{"../benchmarks/parser/type checking/good/good7.golite", []string{""}},

		// Bad programs
		{"../benchmarks/parser/type checking/bad/bad1.golite", []string{"semantic error at line (17), col (4): undeclared variable fir. (undeclared variable)", "semantic error at line (17), col (4): type mismatch in assignment (type mistmatch)", "semantic error at line (21), col (8): variable fir not declared. (undeclared variable)", "semantic error at line (21), col (8): invalid expression (mismatched types)", "semantic error at line (22), col (1): undeclared variable fir. (undeclared variable)", "semantic error at line (22), col (1): type mismatch in assignment (type mistmatch)", "semantic error at line (25), col (11): variable fir not declared. (undeclared variable)", "semantic error at line (25), col (4): return type does not match function signature return type (mismatched types)", "semantic error at line (36), col (8): invalid expression (mismatched types)", "Failed semantic analysis"}},
		{"../benchmarks/parser/type checking/bad/bad2.golite", []string{"semantic error at line (27), col (2): undeclared variable root. (undeclared variable)", "semantic error at line (27), col (2): type mismatch in assignment (type mistmatch)", "semantic error at line (41), col (7): field right not declared in type *Node (undeclared field)", "semantic error at line (44), col (5): field right not declared in type *Node (undeclared field)", "semantic error at line (46), col (23): field right not declared in type *Node (undeclared field)", "semantic error at line (61), col (5): field right not declared in type *Node (undeclared field)", "semantic error at line (62), col (18): field right not declared in type *Node (undeclared field)", "semantic error at line (72), col (5): field right not declared in type *Node (undeclared field)", "semantic error at line (73), col (20): field right not declared in type *Node (undeclared field)", "semantic error at line (82), col (1): undeclared variable root. (undeclared variable)", "semantic error at line (82), col (1): type mismatch in assignment (type mistmatch)", "semantic error at line (83), col (1): type mismatch in assignment (type mistmatch)", "semantic error at line (88), col (18): variable root not declared. (undeclared variable)", "semantic error at line (88), col (2): argument type mismatch in function call (mismatched type)", "semantic error at line (92), col (16): variable root not declared. (undeclared variable)", "semantic error at line (92), col (1): argument type mismatch in function call (mismatched type)", "semantic error at line (93), col (18): variable root not declared. (undeclared variable)", "semantic error at line (93), col (1): argument type mismatch in function call (mismatched type)", "Failed semantic analysis"}},
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

	// Remove the IR created from the tests
	utils.RemoveFolder("IR")
	// Remove the assembly folder created from the tests
	utils.RemoveFolder("assembly")
}
