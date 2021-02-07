package main

import "fmt"

func sumDigits(n int) int {
	if n == 0 {
		return 0
	}
	return n%10 + sumDigits(n/10)
}

func main() {
	fmt.Println(sumDigits(9990))
}
