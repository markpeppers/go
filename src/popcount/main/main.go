package main

import (
	"fmt"
	"log"
	"os"
	"popcount"
	"strconv"
)

func main() {
	num, err := strconv.ParseUint(os.Args[1], 10, 0)
	if err != nil {
		log.Fatalf("Bad input?: %v", err)
	}
	count := popcount.PopCount(num)
	fmt.Printf("%s: %d\n", os.Args[1], count)
	loopcount := popcount.LoopPopCount(num)
	fmt.Printf("Loop %s: %d\n", os.Args[1], loopcount)
	shiftcount := popcount.ShiftPopCount(num)
	fmt.Printf("Shift %s: %d\n", os.Args[1], shiftcount)
	clearcount := popcount.ClearPopCount(num)
	fmt.Printf("Clear %s: %d\n", os.Args[1], clearcount)
}
