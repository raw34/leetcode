package _0368

import "sort"

func largestDivisibleSubset(nums []int) []int {
    sort.Ints(nums)
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }

    n := len(nums)
    dp := make([]int, n)
    dp[0] = 1
    maxLen := 1
    maxNum := nums[0]
    for i := 1; i < n; i++ {
        dp[i] = 1
        for j := 0; j < i; j++ {
            if nums[i]%nums[j] == 0 {
                dp[i] = max(dp[i], dp[j]+1)
            }
        }
        if dp[i] > maxLen {
            maxLen = dp[i]
            maxNum = nums[i]
        }
    }

    res := make([]int, 0)
    for i := n - 1; i >= 0; i-- {
        if dp[i] == maxLen && maxNum%nums[i] == 0 {
            maxNum = nums[i]
            maxLen--
            res = append([]int{maxNum}, res...)
        }
    }

    return res
}
