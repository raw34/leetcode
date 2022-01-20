package _0279

func numSquares(n int) int {
    min := func(a, b int) int {
        if a > b {
            return b
        }
        return a
    }
    dp := make([]int, n+1)
    for i := 1; i < n+1; i++ {
        dp[i] = i
        for j := 1; i-j*j >= 0; j++ {
            dp[i] = min(dp[i], dp[i-j*j]+1)
        }
    }

    return dp[n]
}
