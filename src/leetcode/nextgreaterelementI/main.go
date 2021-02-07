package main

import "fmt"

func main() {
	nums1 := []int{2, 4}
	nums2 := []int{1, 2, 3, 4}
	fmt.Println(nextGreaterElement(nums1, nums2))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ans := make([]int, len(nums1))
	for i, num1 := range nums1 {
		rem := remaining(nums2, num1)
		ans[i] = searchGreater(rem, num1)
	}
	return ans
}

func remaining(nums []int, search int) []int {
	for i, n := range nums {
		if n == search {
			return nums[i+1:]
		}
	}
	return nil
}

func searchGreater(nums []int, search int) int {
	for _, n := range nums {
		if n > search {
			return n
		}
	}
	return -1
}
