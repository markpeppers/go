package main

import (
	"fmt"

	intset "gopl/ex6.1"
)

func main() {
	var x intset.IntSet

	x.Add(1)
	x.Add(2)
	x.Add(64)
	// x.Add(234323)
	// x.Add(2343534332)
	fmt.Println(x.Has(63), x.Has(5))
	fmt.Println("x:", &x)

	y := x.Copy()
	fmt.Println("1) len y:", y.Len())
	y.Remove(2)
	fmt.Println("2) len y:", y.Len())
	fmt.Println(&x)

}
