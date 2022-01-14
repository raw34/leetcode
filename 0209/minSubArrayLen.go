package _0209

import "math"

func minSubArrayLen(target int, nums []int) int {
    min := func(a, b int) int {
        if a > b {
            return b
        }
        return a
    }
    n := len(nums)
    res := math.MaxInt32
    sum := 0
    for l, r := 0, 0; r < n; r++ {
        sum += nums[r]
        for sum >= target {
            res = min(res, r-l+1)
            sum -= nums[l]
            l++
        }
    }
    if res == math.MaxInt32 {
        res = 0
    }

    return res
}
