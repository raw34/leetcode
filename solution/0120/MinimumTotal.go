package _0120

func minimumTotal(triangle [][]int) int {
    min := func(a, b int) int {
        if a < b {
            return a
        } else {
            return b
        }
    }

    n := len(triangle)
    dp := make([][]int, n)
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, n)
    }
    dp[0][0] = triangle[0][0]
    for i := 1; i < n; i++ {
        dp[i][0] = dp[i-1][0] + triangle[i][0]
        for j := 1; j < i; j++ {
            dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
        }
        dp[i][i] = dp[i-1][i-1] + triangle[i][i]
    }

    res := dp[n-1][0]
    for _, v := range dp[n-1] {
        res = min(res, v)
    }

    return res
}
