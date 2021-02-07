package main

import "fmt"

func main() {
	nums := []int{1, 2, 1}
	fmt.Println(nextGreaterElements(nums))
}

func nextGreaterElements(nums []int) []int {
	ans := make([]int, len(nums))
	for i := range nums {
		ans[i] = searchCircular(nums, i)
	}
	return ans
}

func searchCircular(nums []int, index int) int {
	num := nums[index]
	N := len(nums)
	i := (index + 1) % N
	for i != index {
		if nums[i] > num {
			return nums[i]
		}
		i++
		i = i % N
	}
	return -1
}
