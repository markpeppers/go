package main

import "fmt"

const n = 12

func main() {
	nums := make([]int, 0)
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}
	perms := permute(nums)
	/*
			for i := range perms {
				fmt.Println(perms[i])
		    }
	*/
	fmt.Println(len(perms), "permutations")
}

func permute(nums []int) [][]int {
	// Calculate all permutatons of 2 integers, then recurse
	if len(nums) == 1 {
		return [][]int{
			{
				nums[0],
			},
		}
	}
	p := permute(nums[1:])
	plen := len(p)
	n := len(nums)
	out := make([][]int, 0)
	for i := 0; i < plen; i++ {
		for j := 0; j < n; j++ {
			out = append(out, insertAt(p[i], nums[0], j))
		}
	}
	return out
}

func insertAt(in []int, n, pos int) []int {
	out := make([]int, len(in)+1)
	for i := 0; i <= len(in); i++ {
		if i < pos {
			out[i] = in[i]
		} else if i == pos {
			out[i] = n
		} else {
			out[i] = in[i-1]
		}
	}
	return out
}
