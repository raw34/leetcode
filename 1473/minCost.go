package _1473

import (
    "math"
)

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
    min := func(index int, arr ...int) int {
        res := math.MaxInt32
        for i, num := range arr {
            if num < res && i != index {
                res = num
            }
        }

        return res
    }

    // dp[i][j][k] 表示第i个房子在第j个街区刷第k种颜色时的最小花费
    dp := make([][][]int, m)
    // 初始化操作
    for i := 0; i < m; i++ {
        dp[i] = make([][]int, target+1)
        for j := 0; j < target+1; j++ {
            dp[i][j] = make([]int, n+1)
            for k := 0; k < n+1; k++ {
                if i == 0 && j == 1 && k == houses[0] && houses[0] > 0 {
                    // 第1个房子第1个街区的情况下，如果第1个房子已有颜色，花费为0
                    dp[i][j][k] = 0
                } else if i == 0 && j == 1 && k > 0 && houses[0] == 0 {
                    // 第1个房子第1个街区的情况下，如果第1个房子没有颜色，花费为第1个房子当前颜色的金额
                    dp[i][j][k] = cost[i][k-1]
                } else {
                    // 其他情况全都初始化为不可用金额
                    dp[i][j][k] = math.MaxInt32
                }
            }
        }
    }

    for i := 1; i < m; i++ {
        for j := 1; j < target+1; j++ {
            k := houses[i]
            if houses[i] > 0 {
                // 当前房子已有颜色
                // 获取变街区和不变街区的花费，取最小值
                dp[i][j][k] = min(-1, dp[i-1][j][k], min(k, dp[i-1][j-1]...))
            } else {
                // 当前房子没有颜色，遍历每种颜色，分别计算花费
                for k = 1; k < n+1; k++ {
                    // 获取变街区和不变街区的花费，取最小值
                    dp[i][j][k] = min(-1, dp[i-1][j][k], min(k, dp[i-1][j-1]...)) + cost[i][k-1]
                }
            }
        }
    }

    res := min(-1, dp[m-1][target]...)
    if res == math.MaxInt32 {
        res = -1
    }

    return res
}
