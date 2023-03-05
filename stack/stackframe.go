package stack

import (
	"sort"
	"strconv"
	"strings"
)

type StackFrame struct {
	// The name of the stack table (corresponds to function name)
	name string
	// The stack table
	table map[string]string
}

// Create a new stack table
func NewStackFrame(name string) *StackFrame {
	return &StackFrame{name, make(map[string]string)}
}

// Add a new entry to the stack table
func (s *StackFrame) AddEntry(name string, location string) {
	s.table[name] = location
}

// String representation of the stack table
func (s *StackFrame) String() string {
	var out string
	out += "Stack Frame: " + s.name + "\n"
	regs := make([][]string, 0)
	rest := make([][]string, 0)
	for name, location := range s.table {
		if strings.Contains(location, "x") {
			regs = append(regs, []string{name, location})
		} else {
			rest = append(rest, []string{name, location})
		}
	}
	// Sort the registers and the rest of the variables by their offset
	// registers in ascending order
	// rest in descending order
	sort.Slice(regs, func(i, j int) bool {
		r1 := regs[i][1]
		r1Priority, _ := strconv.Atoi(r1[1:])

		r2 := regs[j][1]
		r2Priority, _ := strconv.Atoi(r2[1:])
		return r1Priority < r2Priority
	})

	sort.Slice(rest, func(i, j int) bool {
		r1 := rest[i][1]
		r1Offset, _ := strconv.Atoi(r1)

		r2 := rest[j][1]
		r2Offset, _ := strconv.Atoi(r2)
		return r1Offset > r2Offset
	})

	for _, pair := range regs {
		name := pair[0]
		location := pair[1]
		out += name + " -> " + "sp + " + location + "\n"
	}
	for _, pair := range rest {
		name := pair[0]
		location := pair[1]
		out += name + " -> " + "sp + " + location + "\n"
	}
	return out
}

// Get the location of a variable in the stack table
func (s *StackFrame) GetLocation(name string) string {
	return s.table[name]
}

// Get the largest offset in the stack table
func (s *StackFrame) GetLargestOffset() int {
	var largest int
	for name, location := range s.table {
		if !strings.Contains(name, "x") {
			offset, _ := strconv.Atoi(location)
			if offset > largest {
				largest = offset
			}
		}
	}
	return largest
}
