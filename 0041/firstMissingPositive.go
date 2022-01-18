package _0041

import "math"

func firstMissingPositive(nums []int) int {
    hash := map[int]bool{}
    max := nums[0]
    for _, num := range nums {
        hash[num] = true
        if num > max {
            max = num
        }
    }
    res := math.MaxInt32
    for i := 1; i <= max+1; i++ {
        if !hash[i] {
            res = i
            break
        }
    }
    if res == math.MaxInt32 {
        res = 1
    }

    return res
}
