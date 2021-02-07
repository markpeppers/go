package main

import "fmt"

const UintSize = 32 << (^uint(0) >> 32 & 1)
const MaxInt = 1<<(UintSize-1) - 1

func main() {
	C := 17
	coins := []int{1, 2, 5, 10}

	amount := make([]int, C+1)
	for amt := 1; amt <= C; amt++ {
		amount[amt] = MaxInt
		temp := MaxInt
		for c := 0; c < len(coins); c++ {
			if coins[c] <= amt {
				temp_amt := amount[amt-coins[c]] + 1
				if temp_amt < temp {
					temp = temp_amt
					amount[amt] = temp
				}
			}
		}
	}
	fmt.Println(amount)
	fmt.Println(amount[C])
}
