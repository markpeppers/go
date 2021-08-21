package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(camelcase("helloWorldAss"))
	fmt.Println(ceaserCipher("acxyz-ACXYZ?", 3))
}

func camelcase(s string) int32 {
	if len(s) == 0 {
		return 0
	}

	wordcount := int32(1)

	for _, c := range s {
		if string(c) == strings.ToUpper(string(c)) {
			wordcount++
		}
	}
	return wordcount
}

func ceaserCipher(s string, k int32) string {
	result := make([]byte, len(s))
	bytes := []byte(s)
	// var lowerBound byte
	for i, c := range bytes {
		if !((c >= 'a' && c <= 'z') || c >= 'A' && c <= 'Z') {
			result[i] = c
			continue
		}
		result[i] = rotk(c, k)
	}
	return string(result)
}

func rotk(c byte, k int32) byte {
	var firstLetter byte
	if c >= 'a' && c <= 'z' {
		firstLetter = 'a'
	} else {
		firstLetter = 'A'
	}
	return (c-firstLetter+byte(k))%26 + firstLetter
}
