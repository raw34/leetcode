package _0123

func maxProfit(prices []int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }

    n := len(prices)
    dp := make([][3][2]int, n)
    dp[0][1][0] = 0
    dp[0][1][1] = -prices[0]
    dp[0][2][0] = 0
    dp[0][2][1] = -prices[0]
    for i := 1; i < n; i++ {
        dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1]+prices[i])
        dp[i][1][1] = max(dp[i-1][1][1], dp[i-1][0][0]-prices[i])
        dp[i][2][0] = max(dp[i-1][2][0], dp[i-1][2][1]+prices[i])
        dp[i][2][1] = max(dp[i-1][2][1], dp[i-1][1][0]-prices[i])
    }

    return dp[n-1][2][0]
}
