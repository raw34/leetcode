package _0121

func maxProfit(prices []int) int {
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

    res := 0
    n := len(prices)
    minPrice := prices[0]
    for i := 1; i < n; i++ {
        price := prices[i]
        res = max(res, price-minPrice)
        minPrice = min(minPrice, price)
    }

    return res
}

func maxProfit2(prices []int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }

    n := len(prices)
    // dp[i][0] 下标为 i 这天结束的时候，不持股，手上拥有的现金数
    // dp[i][1] 下标为 i 这天结束的时候，持股，手上拥有的现金数
    dp := make([][2]int, n)
    dp[0][0] = 0
    dp[0][1] = -prices[0]
    for i := 1; i < n; i++ {
        dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
        dp[i][1] = max(dp[i-1][1], -prices[i])
    }

    return dp[n-1][0]
}
