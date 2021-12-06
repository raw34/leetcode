package _0322

func coinChange(coins []int, amount int) int {
    min := func(a, b int) int {
        if a < b {
            return a
        }

        return b
    }

    max := amount + 1
    dp := make([]int, max)
    dp[0] = 0
    for i := 1; i <= amount; i++ {
        dp[i] = max
        for _, coin := range coins {
            diff := i - coin
            if diff < 0 {
                continue
            }
            dp[i] = min(dp[i], dp[diff]+1)
        }
    }

    if dp[amount] == max {
        return -1
    }

    return dp[amount]
}
