package main

import (
	"fmt"
	"strings"
)

type token struct {
	s   int
	pos int
}

func balance(s string, toks []token) (string, []token) {
	sum := 0
	for i := 0; i < len(toks); {
		sum += toks[i].s
		if sum < 0 {
			// Remove a ")"
			s = markChar(s, toks[i].pos)
			toks = append(toks[:i], toks[i+1:]...)
			sum += 1
		} else {
			i++
		}
	}
	sum = 0
	for i := len(toks) - 1; i >= 0; {
		sum += toks[i].s
		if sum > 0 {
			// Remove a "("
			s = markChar(s, toks[i].pos)
			toks = append(toks[:i], toks[i+1:]...)
			sum -= 1
		}
		i--
	}
	return s, toks
}

func markChar(s string, idx int) string {
	return s[:idx] + "*" + s[idx+1:]
}

func minRemoveToMakeValid(s string) string {
	toks := make([]token, 0)
	for i, c := range s {
		if string(c) == "(" {
			toks = append(toks, token{1, i})
		}
		if string(c) == ")" {
			toks = append(toks, token{-1, i})
		}
	}
	s, toks = balance(s, toks)
	return strings.Replace(s, "*", "", -1)
}

func main() {
	strings := []string{
		"lee(t(c)o)de)",
		"a)b(c)d",
		"))((",
		"(a(b(c)d)",
		"((())",
		"))a()d(",
	}
	for _, s := range strings {
		fmt.Println(s)
		fmt.Println(minRemoveToMakeValid(s))
		fmt.Println()
	}
}
