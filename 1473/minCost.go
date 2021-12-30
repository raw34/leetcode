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

    dp := make([][][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([][]int, target+1)
        for j := 0; j < target+1; j++ {
            dp[i][j] = make([]int, n+1)
            for k := 0; k < n+1; k++ {
                if i == 0 && j == 1 && k == houses[0] && houses[0] > 0 {
                    dp[i][j][k] = 0
                } else if i == 0 && j == 1 && k > 0 && houses[0] == 0 {
                    dp[i][j][k] = cost[i][k-1]
                } else {
                    dp[i][j][k] = math.MaxInt32
                }
            }
        }
    }

    for i := 1; i < m; i++ {
        for j := 1; j < target+1; j++ {
            color := houses[i]
            if houses[i] > 0 {
                for k := 1; k < n+1; k++ {
                    if k == color {
                        dp[i][j][color] = min(-1, dp[i][j][color], dp[i-1][j][color])
                    } else {
                        dp[i][j][color] = min(-1, dp[i][j][color], dp[i-1][j-1][k])
                    }
                }
            } else {
                for color = 1; color < n+1; color++ {
                    for k := 1; k < n+1; k++ {
                        if k == color {
                            dp[i][j][color] = min(-1, dp[i][j][color], dp[i-1][j][k]+cost[i][color-1])
                        } else {
                            dp[i][j][color] = min(-1, dp[i][j][color], dp[i-1][j-1][k]+cost[i][color-1])
                        }
                    }
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
