package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	index, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err, "Goodbye")
	}
	row := getRow(index)
	fmt.Println(row)
}

func getRow(rowIndex int) []int64 {
	memos := make(map[int][]int64, 0)
	if rowIndex == 0 {
		return []int64{1}
	}
	if rowIndex == 1 {
		return []int64{1, 1}
	}
	row := make([]int64, rowIndex+1)
	for i := 0; i < rowIndex+1; i++ {
		if i == 0 || i == rowIndex {
			row[i] = 1
			continue
		}
		lastRow := getRowHelper(rowIndex-1, memos)
		row[i] = lastRow[i-1] + lastRow[i]
	}
	return row
}

func getRowHelper(rowIndex int, memos map[int][]int64) []int64 {
	r, found := memos[rowIndex]
	if found {
		return r
	}
	memos[rowIndex] = getRow(rowIndex)
	return memos[rowIndex]
}
