package main

import (
	"fmt"
	"sort"
)

func main() {
	A := []int{8038, 9111, 5458, 8483, 5052, 9161, 8368, 2094, 8366, 9164, 53, 7222, 9284, 5059, 4375, 2667, 2243, 5329, 3111, 5678, 5958, 815, 6841, 1377, 2752, 8569, 1483, 9191, 4675, 6230, 1169, 9833, 5366, 502, 1591, 5113, 2706, 8515, 3710, 7272, 1596, 5114, 3620, 2911, 8378, 8012, 4586, 9610, 8361, 1646, 2025, 1323, 5176, 1832, 7321, 1900, 1926, 5518, 8801, 679, 3368, 2086, 7486, 575, 9221, 2993, 421, 1202, 1845, 9767, 4533, 1505, 820, 967, 2811, 5603, 574, 6920, 5493, 9490, 9303, 4648, 281, 2947, 4117, 2848, 7395, 930, 1023, 1439, 8045, 5161, 2315, 5705, 7596, 5854, 1835, 6591, 2553, 8628}
	K := 4643
	// A := []int{1}
	// K := 0
	fmt.Println(smallestRangeII(A, K))
}

func smallestRangeII(A []int, K int) int {
	sort.Sort(sort.IntSlice(A))
	lenA := len(A)
	arr := make([]int, lenA)
	smallestDiff := 10000
	done := false
	for !done {
		// fmt.Println(arr)
		B := make([]int, lenA)
		for j := 0; j < len(A); j++ {
			mult := 1
			if arr[j] == 0 {
				mult = -1
			}
			B[j] = A[j] + mult*K
			if maxMinDiff(B, j+1) > smallestDiff {
				fmt.Println("skipping")
				j = len(A)
			}
		}
		// fmt.Println(B)
		// fmt.Println(maxMinDiff(B))
		if maxMinDiff(B, len(A)) < smallestDiff {
			smallestDiff = maxMinDiff(B, len(A))
		}
		done = inc(arr)
	}
	return smallestDiff
}

func inc(in []int) bool {
	pos := len(in) - 1
	in[pos]++
	for in[pos] > 1 {
		in[pos] = 0
		pos--
		if pos < 0 {
			return true
		}
		in[pos]++
	}
	return false
}

func max(A []int, len int) int {
	m := 0
	for i := 0; i < len; i++ {
		if A[i] > m {
			m = A[i]
		}
	}
	return m
}

func min(A []int, len int) int {
	m := 10000
	for i := 0; i < len; i++ {
		if A[i] < m {
			m = A[i]
		}
	}
	return m
}

func maxMinDiff(A []int, len int) int {
	return max(A, len) - min(A, len)
}
