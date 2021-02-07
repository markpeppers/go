package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("ASS")
		os.Exit(1)
	}
	n, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("FU")
		os.Exit(1)
	}
	rand.Seed(time.Now().UnixNano())
	randPrices := make([]int, n)
	for i := range randPrices {
		randPrices[i] = rand.Intn(15)
	}
	fmt.Println(randPrices)
	switch os.Args[1] {
	case "s":
		fmt.Println(maxPrice(n, randPrices))
	case "f":
		fmt.Println(maxPriceDP(n, randPrices))
	}
}
func maxPrice(n int, prices []int) int {
	if n == 0 {
		return 0
	}
	max := -1
	for i := 0; i < n; i++ {
		tmp := prices[n-i-1] + maxPrice(i, prices)
		if tmp > max {
			max = tmp
		}
	}
	return max
}

func maxPriceDP(n int, prices []int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i < n+1; i++ {
		max := -1
		for j := 1; j <= i; j++ {
			tmp := prices[j-1] + dp[i-j]
			if tmp > max {
				max = tmp
			}
		}
		dp[i] = max
	}
	return dp[n]
}
