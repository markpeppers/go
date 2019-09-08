package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counters := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counters[input.Text()]++
	}
	for line, n := range counters {
		if n > 1 {
			fmt.Printf("%d: %s\n", n, line)
		}
	}
}
