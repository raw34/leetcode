package _0712

func minimumDeleteSum(s1 string, s2 string) int {
    // 可转化为最长公共子序列问题
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }
    m := len(s1)
    n := len(s2)
    // 求所有字符ascii码和
    sum := 0
    for i := 0; i < m; i++ {
        sum += int(s1[i])
    }
    for j := 0; j < n; j++ {
        sum += int(s2[j])
    }

    // 求公共子序列ascii码和
    dp := make([][]int, m+1)
    dp[0] = make([]int, n+1)
    for i := 1; i < m+1; i++ {
        dp[i] = make([]int, n+1)
        for j := 1; j < n+1; j++ {
            if s1[i-1] == s2[j-1] {
                dp[i][j] = dp[i-1][j-1] + int(s1[i-1])
            } else {
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
            }
        }
    }

    // 结果 = 总和 - 双倍公共子序列和
    return sum - dp[m][n]*2
}
