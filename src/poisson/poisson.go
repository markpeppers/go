package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	lambda := 10.0
	L := math.Exp(-lambda)
	k := 0
	p := 1.0
	rand.Seed(time.Now().UnixNano())
	for p > L {
		k = k + 1
		u := rand.Float64()
		p = p * u
	}
	fmt.Println(k - 1)
}
