package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	amt, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Problem")
		os.Exit(1)
	}
	fmt.Println(coinChange(amt, []int{1, 2, 5}))
}
func coinChange(amount int, denominations []int) int {
	combinations := make([]int, amount+1)
	combinations[0] = 1

	for j := 0; j < len(denominations); j++ {
		coin := denominations[j]
		for higherAmount := coin; higherAmount <= amount; higherAmount++ {
			remainder := higherAmount - coin
			// fmt.Printf("combinations[%d] += combinations[%d]\n", higherAmount, remainder)
			combinations[higherAmount] += combinations[remainder]
			// fmt.Println(combinations)
		}
	}
	return combinations[amount]
}
