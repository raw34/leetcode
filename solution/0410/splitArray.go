package _0410

import "math"

func splitArray(nums []int, m int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }
    min := func(a, b int) int {
        if a > b {
            return b
        }
        return a
    }
    n := len(nums)
    preSum := make([]int, n+1)
    for i := 1; i < n+1; i++ {
        preSum[i] = preSum[i-1] + nums[i-1]
    }
    dp := make([][]int, n+1)
    for i := 0; i < n+1; i++ {
        dp[i] = make([]int, m+1)
        for j := 0; j < m+1; j++ {
            dp[i][j] = math.MaxInt32
        }
    }
    dp[0][0] = 0
    for i := 1; i < n+1; i++ {
        for j := 1; j < m+1; j++ {
            for k := 0; k < i; k++ {
                dp[i][j] = min(dp[i][j], max(dp[k][j-1], preSum[i]-preSum[k]))
            }
        }
    }

    return dp[n][m]
}
