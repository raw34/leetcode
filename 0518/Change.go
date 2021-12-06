package _0518

func change(amount int, coins []int) int {
    n := amount + 1
    dp := make([]int, n)
    dp[0] = 1

    for _, coin := range coins {
        for i := coin; i < n; i++ {
            diff := i - coin
            dp[i] += dp[diff]
        }
    }

    return dp[amount]
}
