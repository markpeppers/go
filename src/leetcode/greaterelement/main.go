package main

import "fmt"

func main() {
	n := 12343232423559
	fmt.Println(nextGreaterElement(n))
	fmt.Println(n)
	fmt.Println(decompose(n))
}

func nextGreaterElement(n int) int {
	return 0
}

func decompose(n int) []int {
	digits := make([]int, 0)
	for n > 10 {
		digits = append(digits, n%10)
		n = n / 10
	}
	digits = append(digits, n)
	return reverse(digits)
}

func reverse(numbers []int) []int {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}
