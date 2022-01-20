package _0045

import "math"

func jump(nums []int) int {
    min := func(a, b int) int {
        if a < b {
            return a
        }
        return b
    }

    n := len(nums)
    dp := make([]int, n)
    dp[0] = 0
    for i := 1; i < n; i++ {
        jumps := math.MaxInt32
        for j := 0; j < i; j++ {
            step := i - j
            if nums[j] >= step {
                jumps = min(dp[j]+1, jumps)
            }
        }
        dp[i] = jumps
    }

    return dp[n-1]
}
