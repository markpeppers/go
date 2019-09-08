package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	var output, sep string
	for _, s := range(os.Args[1:]) {
		output += sep + s
		sep = " "
	}
	fmt.Println(output)
	nanos := time.Since(start).Nanoseconds()
	fmt.Println(nanos)
}
