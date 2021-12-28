package _0091

func numDecodings(s string) int {
    n := len(s)
    dp := make([]int, n+1)
    dp[0] = 1
    for i := 1; i <= n; i++ {
        // 前一位数字不为0，初始化当前位解法数为前一位解法数
        if s[i-1] != '0' {
            dp[i] = dp[i-1]
        }
        // 当前一位为两位数时，继续累加前前一位的解法数
        if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+(s[i-1]-'0')) <= 26 {
            dp[i] += dp[i-2]
        }
    }

    return dp[n]
}
