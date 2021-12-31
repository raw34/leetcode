package _0115

func numDistinct(s string, t string) int {
    m := len(t)
    n := len(s)
    // dp[i][j]表示t串前i个字符可有s串前j个字符组成的个数
    dp := make([][]int, m+1)
    for i := 0; i < m+1; i++ {
        dp[i] = make([]int, n+1)
    }
    // 初始化第一列，当t串取空是，任意长s串都可以组成一个t串
    for j := 0; j < n+1; j++ {
        dp[0][j] = 1
    }

    for i := 1; i < m+1; i++ {
        for j := i; j < n+1; j++ {
            if t[i-1] == s[j-1] {
                // 当前t和当前s相同时，取t串前移一位的情况 + t串和s串同时前移一位的情况
                dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
            } else {
                // 当前t和当前s不同时，取t串前移一位的情况
                dp[i][j] = dp[i][j-1]
            }
        }
    }

    return dp[m][n]
}
