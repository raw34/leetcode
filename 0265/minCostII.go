package _0265

import "math"

func minCostII(costs [][]int) int {
    min := func(index int, arr ...int) int {
        res := math.MaxInt32
        for i, num := range arr {
            if num < res && i != index {
                res = num
            }
        }

        return res
    }

    m := len(costs)
    n := len(costs[0])
    // dp[i][j]表示第i个房子刷第j种颜色的情况下的最少花费
    dp := make([][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
    }
    // 初始化第1个房子
    for j := 0; j < n; j++ {
        dp[0][j] = costs[0][j]
    }
    // 遍历每个房子
    for i := 1; i < m; i++ {
        // 当前房子刷当前颜色的总花费 = 前一个房子的不刷当前颜色时的最少消费 + 当前房子当前颜色消费
        for j := 0; j < n; j++ {
            dp[i][j] = min(j, dp[i-1]...) + costs[i][j]
        }
    }

    return min(-1, dp[m-1]...)
}
