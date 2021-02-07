package threesum

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		for j := i + 1; j < len(nums); j++ {
			sum2 := nums[i] + nums[j]
			if sum2 > 0 {
				break
			}
			for k := j + 1; k < len(nums); k++ {
				if sum2+nums[k] == 0 {
					if isFound(res, nums[i], nums[j]) {
						break
					}
					res = append(res, []int{nums[i], nums[j], nums[k]})
					// No number beyond k is going to satisfy
					break
				}
				if sum2+nums[k] > 0 {
					break
				}
			}
		}
	}
	return res
}

func isFound(res [][]int, n1, n2 int) bool {
	for _, row := range res {
		if n1 == row[0] && n2 == row[1] {
			return true
		}
	}
	return false
}
