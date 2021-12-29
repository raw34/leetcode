package _0122

func maxProfit(prices []int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }

    n := len(prices)
    // dp[i][0] 下标为 i 这天结束的时候，不持股，可获得的最大收益
    // dp[i][1] 下标为 i 这天结束的时候，持股，可获得的最大收益
    dp := make([][2]int, n)
    dp[0][0] = 0
    dp[0][1] = -prices[0]
    for i := 1; i < n; i++ {
        // 当天不持股的情况下，获取前一天不持股和前一天持股今天卖出的最大收益
        dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
        // 当天持股的情况下，获取前一天持股和前一天不持股今天买入的最大收益
        dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
    }

    return dp[n-1][0]
}

func maxProfit2(prices []int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }

    n := len(prices)
    cash, preCash := 0, 0
    hold, preHold := -prices[0], -prices[0]
    for i := 1; i < n; i++ {
        cash = max(preCash, hold+prices[i])
        hold = max(preHold, cash-prices[i])

        preCash = cash
        preHold = hold
    }

    return cash
}
