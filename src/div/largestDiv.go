func largestDivisibleSubset(nums []int) []int {
    sort.Ints(nums)
    return largestDivisibleSubsetHelper(nums)
}

func largestDivisibleSubsetHelper(nums [] int) []int {   
    if len(nums) == 0 {
        return []int{}
    }
    if len(nums) == 1 {
        return nums
    }
    if isDivisible(nums) {
        return nums
    }
    subs := subsets(nums)
    candidates := make([][]int, 0)
    for _, subset := range subs {
        foo := largestDivisibleSubset(subset)
        candidates = append(candidates, foo)
    }
    largestSub := make([]int, 0)
    for _, c := range candidates {
        if len(c) > len(largestSub) {
            largestSub = c
        }
    }
    return largestSub
}

func subsets(nums []int) [][]int {
    sets := make([][]int, len(nums))
    for i := 0; i < len(nums); i++ {
        sets[i] = make([]int, 0, len(nums) - 1)
        for j := 0; j < len(nums); j++ {
            if j == i {
                continue
            }
            sets[i] = append(sets[i], nums[j])
        }
    }
    return sets
}

func isDivisible(nums []int) bool {
    for i := range(nums) {
        if nums[i] == 1 {
            continue    // speedup
        }
        for j := i + 1; j < len(nums); j++ {
            if nums[j] % nums[i] != 0 {
                return false
            }
        }
    }
    return true
}
