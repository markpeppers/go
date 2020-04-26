package main

import (
	"fmt"
	"tempconv"
)

func main() {
	fmt.Printf("%s is %s", tempconv.AbsoluteZeroC, tempconv.CtoF(AbsoluteZeroC))
}
