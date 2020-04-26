package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	vals, err := Matoi(os.Args[1:]...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "input error: %v", err)
	}
	fmt.Println("Min:", min(vals...))
	os.Exit(1)
}

func min(vals ...int) int {
	if len(vals) == 0 {
		panic("No minimum")
	}
	ret := math.MaxInt64
	for _, val := range vals {
		if val < ret {
			ret = val
		}
	}
	return ret
}

// Matoi for Multiple Atoi converts an array of strings to ints
func Matoi(args ...string) ([]int, error) {
	ints := make([]int, 0, len(args))
	for _, arg := range args {
		val, err := strconv.Atoi(arg)
		if err != nil {
			return nil, err
		}
		ints = append(ints, val)
	}
	return ints, nil
}
