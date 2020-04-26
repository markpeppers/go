package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func expand(s string, f func(string) string) string {
	var out []byte
	for i := 0; i < len(s); i++ {
		if s[i] == '$' {
			end := strings.IndexFunc(s[i+1:], func(r rune) bool {
				return !unicode.IsLetter(r) && !unicode.IsNumber(r)
			})
			if end < 1 {
				end = len(s[i+1:]) + 1
			}
			if i+end == len(s) {
				end--
			}
			out = append(out, []byte(f((s[i+1 : i+end+1])))...)
			i += end + 1
			if i < len(s) {
				out = append(out, s[i])
			}
			continue
		}
		out = append(out, byte(s[i]))
	}
	return string(out)
}

func main() {
	for _, s := range os.Args[1:] {
		fmt.Println(expand(s, replace))
	}
}

func replace(s string) string {
	r := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	if n, found := r[s]; found {
		return n
	}
	return "X"
}
