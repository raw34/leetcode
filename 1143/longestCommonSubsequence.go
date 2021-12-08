package _1143

func longestCommonSubsequence(text1 string, text2 string) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }

    m := len(text1)
    n := len(text2)
    dp := make([][]int, m+1)
    for i := 0; i <= m; i++ {
        dp[i] = make([]int, n+1)
    }

    for i := 1; i <= m; i++ {
        c1 := text1[i-1]
        for j := 1; j <= n; j++ {
            c2 := text2[j-1]
            if c1 == c2 {
                dp[i][j] = dp[i-1][j-1] + 1
            } else {
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
            }
        }
    }

    return dp[m][n]
}
