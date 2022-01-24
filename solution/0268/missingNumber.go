package _0268

import "sort"

func missingNumber(nums []int) int {
    res := len(nums)
    for i, num := range nums {
        res ^= i ^ num
    }
    return res
}

func missingNumber2(nums []int) int {
    sort.Ints(nums)
    n := len(nums)
    hash := map[int]bool{}
    for i := 0; i < n; i++ {
        num := nums[i]
        if num > 0 && !hash[num-1] {
            return num - 1
        }
        hash[num] = true
    }

    return n
}
