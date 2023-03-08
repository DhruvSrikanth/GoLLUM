package utils

import "fmt"

// Check if a slice contains a string
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// Function that converts an integer to a string represented in 64 bits
func IntToHex(n int64) string {
	return fmt.Sprintf("%016x", n)
}

// Split hex string into 4 chunks representing 16 bits
func SplitHex(hex string) []string {
	var result []string

	for i := 0; i < len(hex); i += 4 {
		result = append(result, "0x"+hex[i:i+4])
	}

	return result
}
