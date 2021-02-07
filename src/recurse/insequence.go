package main

import "fmt"

func inSequence(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	/*
		if len(nums) == 2 {
			if nums[1] == nums[0]+1 {
				return true
			}
			return false
		}
	*/
	return nums[len(nums)-1] == nums[len(nums)-2]+1 && inSequence(nums[:len(nums)-1])
}

func main() {
	nums := []int{3, 3, 5, 6, 7}
	fmt.Println(inSequence(nums))
}
