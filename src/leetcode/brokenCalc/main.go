package main

import "fmt"

func main() {
	fmt.Println(brokenCalc(1000000000, 1))
}

func brokenCalc(X int, Y int) int {
	inc := 0
	for Y > X {
		if Y%2 == 1 {
			Y++
		} else {
			Y /= 2
		}
		inc++
	}
	return inc + X - Y
}

func brokenCalcSLOW(X int, Y int) int {
	inc := 0
	for X != Y {
		if Y%2 == 1 || X > Y {
			Y++
		} else {
			Y /= 2
		}
		inc++
	}
	return inc
}
