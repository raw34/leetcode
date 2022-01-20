package _0087

func isScramble(s1 string, s2 string) bool {
    if s1 == s2 {
        return true
    }
    if len(s1) != len(s2) {
        return false
    }

    n := len(s1)
    dp := make([][][]bool, n)
    for i := 0; i < n; i++ {
        dp[i] = make([][]bool, n)
        for j := 0; j < n; j++ {
            dp[i][j] = make([]bool, n+1)
            dp[i][j][1] = s1[i] == s2[j]
        }
    }
    for l := 2; l <= n; l++ {
        for i := 0; i <= n-l; i++ {
            for j := 0; j <= n-l; j++ {
                for k := 1; k <= l-1; k++ {
                    a := dp[i][j][k] && dp[i+k][j+k][l-k]
                    b := dp[i][j+l-k][k] && dp[i+k][j][l-k]
                    if a || b {
                        dp[i][j][l] = true
                    }
                }
            }
        }
    }

    return dp[0][0][n]
}
