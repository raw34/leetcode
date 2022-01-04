package _0174

import "math"

func calculateMinimumHP(dungeon [][]int) int {
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

    m := len(dungeon)
    n := len(dungeon[0])
    dp := make([][]int, m+1)
    for i := 0; i < m+1; i++ {
        dp[i] = make([]int, n+1)
        for j := 0; j < n+1; j++ {
            dp[i][j] = math.MaxInt32
        }
    }
    dp[m-1][n] = 1
    dp[m][n-1] = 1
    for i := m - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            minn := min(dp[i+1][j], dp[i][j+1])
            dp[i][j] = max(minn-dungeon[i][j], 1)
        }
    }

    return dp[0][0]
}
