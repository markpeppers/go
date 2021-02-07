package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "9"
	for i := 0; i < 10; i++ {
		fmt.Println(s)
		s = looksay(s)
	}

}

func looksay(in string) string {
	out := ""
	for len(in) > 0 {
		n, numStr := gather(in)
		out += numStr + n
		num, _ := strconv.Atoi(numStr)
		in = in[num:]
	}
	return out
}

func gather(in string) (nString, numString string) {
	num := 0
	firstChar := in[0:1]
	for num < len(in) && in[num:num+1] == firstChar {
		num++
	}
	numString = strconv.Itoa(num)
	nString = firstChar
	return
}
