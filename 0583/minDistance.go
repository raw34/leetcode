package _0583

func minDistance(word1 string, word2 string) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }
    m := len(word1)
    n := len(word2)
    // dp[i][j]表示从两字符串最长公共子序列的长度
    dp := make([][]int, m+1)
    for i := 0; i <= m; i++ {
        dp[i] = make([]int, n+1)
    }
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if word1[i-1] == word2[j-1] {
                // 当前字符相同时，公共子序列长度加1
                dp[i][j] = dp[i-1][j-1] + 1
            } else {
                // 当前字符不同时，取前一个位置最长子串长度
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
            }
        }
    }

    // 结果为两串长度和减两个公共子序列长度
    return m + n - dp[m][n]*2
}
