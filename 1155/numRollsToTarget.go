package _1155

func numRollsToTarget(n int, f int, target int) int {
    mod := int(1e9 + 7)
    dp := make([][]int, n+1)
    for i := 0; i < n+1; i++ {
        dp[i] = make([]int, target+1)
    }
    dp[0][0] = 1
    for i := 1; i < n+1; i++ {
        for j := 1; j < target+1; j++ {
            for k := 1; k < f+1; k++ {
                if j-k >= 0 {
                    dp[i][j] = (dp[i][j] + dp[i-1][j-k]) % mod
                }
            }
        }
    }

    return dp[n][target]
}
