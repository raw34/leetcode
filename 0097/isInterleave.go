package _0097

func isInterleave(s1 string, s2 string, s3 string) bool {
    m := len(s1)
    n := len(s2)
    if m+n != len(s3) {
        return false
    }

    dp := make([][]bool, m+1)
    for i := 0; i < m+1; i++ {
        dp[i] = make([]bool, n+1)
    }
    dp[0][0] = true
    for i := 0; i < m+1; i++ {
        for j := 0; j < n+1; j++ {
            k := i + j
            //a := s1[i-1]
            //b := s2[j-1]
            //c1 := s3[k-1]
            //c2 := s3[k]
            //fmt.Println(i, j, k, string(a), string(b), string(c1), string(c2))
            //if a == c2 && b == c1 || a == c1 && b == c2 {
            //    dp[i][j] = dp[i-1][j-1]
            //} else if a != c1 && a != c2 && (b == c1 || b == c2) {
            //    dp[i][j] = dp[i-1][j]
            //} else if b != c1 && b != c2 && (a == c1 || a == c2) {
            //    dp[i][j] = dp[i][j-1]
            //} else {
            //    dp[i][j] = dp[i][j-1] || dp[i-1][j] || dp[i-1][j-1]
            //}
            if i > 0 {
                dp[i][j] = dp[i][j] || dp[i-1][j] && s1[i-1] == s3[k-1]
            }
            if j > 0 {
                dp[i][j] = dp[i][j] || dp[i][j-1] && s2[j-1] == s3[k-1]
            }
        }
    }

    return dp[m][n]
}
