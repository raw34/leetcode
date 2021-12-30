package _0010

func isMatch(s string, p string) bool {
    m := len(s)
    n := len(p)
    dp := make([][]bool, m+1)
    for i := 0; i < m+1; i++ {
        dp[i] = make([]bool, n+1)
    }
    // 初始化第一个位置
    dp[0][0] = true
    // 初始化第一列，如果遇到*且前一个位置匹配， 则下一个位置赋值为匹配
    for j := 0; j < n; j++ {
        if p[j] == '*' && dp[0][j-1] {
            dp[0][j+1] = true
        }
    }
    for i := 1; i < m+1; i++ {
        for j := 1; j < n+1; j++ {
            if s[i-1] == p[j-1] || p[j-1] == '.' {
                dp[i][j] = dp[i-1][j-1]
            } else if p[j-1] == '*' && s[i-1] != p[j-2] && p[j-2] != '.' {
                dp[i][j] = dp[i][j-2]
            } else if p[j-1] == '*' {
                dp[i][j] = dp[i][j-1] || dp[i][j-2] || dp[i-1][j]
            }
        }
    }

    return dp[m][n]
}
