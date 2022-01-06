package _1049

func lastStoneWeightII(stones []int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }

    n := len(stones)
    // 计算背包大小
    sum := 0
    for i := 0; i < n; i++ {
        sum += stones[i]
    }
    t := sum / 2

    // dp[i][j]表示从前i块石头中选取，选取值之和小于等于目标值j的最大值
    dp := make([][]int, n+1)
    for i := 0; i < n+1; i++ {
        dp[i] = make([]int, t+1)
    }
    for i := 1; i < n+1; i++ {
        stone := stones[i-1]
        for j := 0; j < t+1; j++ {
            if j < stone {
                dp[i][j] = dp[i-1][j]
            } else {
                dp[i][j] = max(dp[i][j], dp[i-1][j-stone]+stone)
            }
        }
    }

    return sum - dp[n][t]*2
}
