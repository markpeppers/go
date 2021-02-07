package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	numerator, _ := strconv.Atoi(os.Args[1])
	denominator, _ := strconv.Atoi(os.Args[2])
	fmt.Println(fractionToDecimal(numerator, denominator))
}

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 && denominator == 0 {
		return "Inf"
	}

	if numerator == 0 {
		return "0"
	}

	var builder strings.Builder

	// Better than ^
	if (numerator < 0) != (denominator < 0) {
		fmt.Fprintf(&builder, "-")
	}

	if denominator == 0 {
		fmt.Fprintf(&builder, "Inf")
		return builder.String()
	}

	// Handle division
	dividend := abs(int64(numerator))
	divisor := abs(int64(denominator))
	fmt.Fprintf(&builder, "%d", dividend/divisor)

	// Return if there is no fraction
	remainder := dividend % divisor
	if remainder == 0 {
		return builder.String()
	}

	// Handle fraction
	fmt.Fprintf(&builder, ".")
	m := make(map[int64]int, 0)
	for remainder != 0 {
		fmt.Println(remainder)
		if pos, exists := m[remainder]; exists {
			fraction := fmt.Sprintf("%s(%s)", builder.String()[0:pos], builder.String()[pos:])
			return fraction
		}

		fmt.Println(builder.String())
		m[remainder] = builder.Len()
		remainder *= 10
		fmt.Fprintf(&builder, "%d", remainder/divisor)
		remainder %= divisor
	}

	return builder.String()
}

func abs(v int64) int64 {
	if v < 0 {
		return -v
	}

	return v
}
