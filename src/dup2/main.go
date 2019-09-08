package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	counts := make(map[string]int)
	containerFiles := make(map[string][]string)
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprint(os.Stderr, "dup2: %s\n", err)
				continue
			}
			countLinesFiles(f, counts, file, containerFiles)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			var sep string
			fmt.Printf("\t")
			for _, filename := range containerFiles[line] {
				fmt.Printf("%s%s", sep, filename)
				sep = ", "
			}
			fmt.Printf("\n")
		}
	}
}

func countLines(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func countLinesFiles(f *os.File, counts map[string]int, filename string, containers map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		//		if containers[input.Text()] == nil {
		//			containers[input.Text()] = make([]string)
		//		}
		if !contains(containers[input.Text()], filename) {
			containers[input.Text()] = append(containers[input.Text()], filename)
		}
	}
}

func contains(array []string, text string) bool {
	for _, s := range array {
		if s == text {
			return true
		}
	}
	return false
}
