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

type stack []int

func (s stack) hasElements() bool {
	return len(s) > 0
}

func (s *stack) pop() int {
	ret := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return ret
}

func (s *stack) push(i int) {
	*s = append(*s, i)
}

func minRemoveToMakeValid2(s string) string {
	stk := stack{}
	sb := []rune(s)
	for i, c := range sb {
		if c == '(' {
			stk.push(i)
		}
		if c == ')' {
			if stk.hasElements() {
				_ = stk.pop()
			} else {
				sb[i] = '*'
			}
		}
	}
	for stk.hasElements() {
		sb[stk.pop()] = '*'
	}
	return strings.ReplaceAll(string(sb), "*", "")
}

func main() {
	strings := []string{
		"lee(t(c)o)de)",
		"a)b(c)d",
		"))((",
		"(a(b(c)d)",
		"((())",
		"))a()d(",
		"",
	}
	for _, s := range strings {
		fmt.Println(s)
		fmt.Println(minRemoveToMakeValid2(s))
		fmt.Println()
	}
}
