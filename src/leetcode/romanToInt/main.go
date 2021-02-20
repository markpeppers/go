package main

import (
	"fmt"
)

func romanToInt(s string) int {
	m := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	sb := []byte(s)
	sum := 0
	curInc := m[sb[len(sb)-1]]
	for i := len(sb) - 1; i >= 0; i-- {
		inc := m[sb[i]]
		if inc < curInc {
			sum -= inc
		} else {
			sum += inc
		}
		curInc = inc
	}
	return sum
}

func main() {
	tests := []struct {
		input    string
		expected int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
		{"MM", 2000},
		{"MMXXI", 2021},
		{"MMMCMXCIX", 3999},
	}
	for _, t := range tests {
		got := romanToInt(t.input)
		check := "\u2713"
		if got != t.expected {
			check = "X"
		}
		fmt.Printf("%9s expected: %5d, got: %5d %s\n", t.input, t.expected, got, check)
	}
}
