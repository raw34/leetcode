package _0714

func maxProfit(prices []int, fee int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }
    n := len(prices)
    dp := make([][2]int, n)
    dp[0][0] = 0
    dp[0][1] = -prices[0]
    for i := 1; i < n; i++ {
        dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
        dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
    }

    return dp[n-1][0]
}
