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

		// Bad programs
		{"../benchmarks/bad/bad1.golite", []string{"func", "main", "(", ")", "{", "var", "a", "bool", ";", "var", "x", "double", ";", "a", "=", "+", ";", "printf", "(", "\"%d\"", ",", "a", ")", ";", "}", "<EOF>", "lexer error(4:9): token recognition error at: '_'"}},
	}
	for num, test := range lexTests {
		testname := fmt.Sprintf("T=%v", num)
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
	}
	for num, test := range parseTests {
		testname := fmt.Sprintf("T=%v", num)
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
