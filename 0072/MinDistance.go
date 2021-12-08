package _0072

func minDistance(word1 string, word2 string) int {
    min := func(a, b, c int) int {
        min := a
        if b < min {
            min = b
        }
        if c < min {
            min = c
        }
        return min
    }

    n1 := len(word1)
    n2 := len(word2)

    dp := make([][]int, n1+1)
    for i := 0; i <= n1; i++ {
        dp[i] = make([]int, n2+1)
    }

    for i := 1; i <= n1; i++ {
        dp[i][0] = dp[i-1][0] + 1
    }

    for j := 1; j <= n2; j++ {
        dp[0][j] = dp[0][j-1] + 1
    }

    for i := 1; i <= n1; i++ {
        for j := 1; j <= n2; j++ {
            if word1[i-1] == word2[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
            }
        }
    }

    return dp[n1][n2]
}
