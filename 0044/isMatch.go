package _0044

func isMatch(s string, p string) bool {
    m := len(s)
    n := len(p)
    dp := make([][]bool, m+1)
    for i := 0; i < m+1; i++ {
        dp[i] = make([]bool, n+1)
    }

    // 初始化第一个位置，两串都不取的情况
    dp[0][0] = true
    // 初始化前缀若干个*的情况
    for j := 1; j < n+1; j++ {
        if p[j-1] != '*' {
            break
        }
        dp[0][j] = true
    }
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            cs := s[i-1]
            cp := p[j-1]
            if cp == '*' {
                dp[i][j] = dp[i][j-1] || dp[i-1][j]
            } else if cs == cp || cp == '?' {
                dp[i][j] = dp[i-1][j-1]
            }
        }
    }

    return dp[m][n]
}
