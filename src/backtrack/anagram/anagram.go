package main

import "fmt"

func main() {
	anagrams("mark")
}

func anagrams(s string) {
	partial := make([]byte, 0)
	used := make([]bool, len(s))
	helper(s, partial, used)
}

func helper(in string, partial []byte, used []bool) {
	if len(partial) == len(in) {
		fmt.Println(string(partial))
		return
	}
	for i := 0; i < len(in); i++ {
		if !used[i] {
			used[i] = true
			partial = append(partial, in[i])
			helper(in, partial, used)
			used[i] = false
			partial = partial[:len(partial)-1]
		}
	}
}
