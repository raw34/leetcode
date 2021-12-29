package _0188

func maxProfit(k int, prices []int) int {
    n := len(prices)
    if n == 0 {
        return 0
    }

    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }

    dp := make([][][2]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([][2]int, k+1)
        for j := 1; j < k+1; j++ {
            dp[0][j][0] = 0
            dp[0][j][1] = -prices[0]
        }
    }
    for i := 1; i < n; i++ {
        for j := 1; j <= k; j++ {
            dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
            dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
        }
    }

    return dp[n-1][k][0]
}
