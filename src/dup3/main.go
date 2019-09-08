package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	counts := make(map[string]int)

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprint(os.Stderr, "dup2: %s\n", err)
			continue
		}
		countLines(f, counts)
		f.Close()
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		counts[input.Text()]++
	}
}
