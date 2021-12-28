package _0256

func minCost(costs [][]int) int {
    min := func(arr ...int) int {
        res := arr[0]
        for _, num := range arr {
            if num < res {
                res = num
            }
        }
        return res
    }

    n := len(costs)
    dp := make([][3]int, n)
    dp[0][0] = costs[0][0]
    dp[0][1] = costs[0][1]
    dp[0][2] = costs[0][2]
    for i := 1; i < n; i++ {
        dp[i][0] = min(dp[i-1][1], dp[i-1][2]) + costs[i][0]
        dp[i][1] = min(dp[i-1][0], dp[i-1][2]) + costs[i][1]
        dp[i][2] = min(dp[i-1][0], dp[i-1][1]) + costs[i][2]
    }

    return min(dp[n-1][0], dp[n-1][1], dp[n-1][2])
}
