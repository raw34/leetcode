package _0097

func isInterleave(s1 string, s2 string, s3 string) bool {
    m := len(s1)
    n := len(s2)
    if m+n != len(s3) {
        return false
    }

    // dp[i][j]表示s1前i个元素和s2前j个元素能否交错组成s3前i+j个元素
    dp := make([][]bool, m+1)
    for i := 0; i < m+1; i++ {
        dp[i] = make([]bool, n+1)
    }
    // 初始化s1和s2都不取的情况
    dp[0][0] = true
    // 初始化第一列，当s3前i个字符都由s1组成的情况
    for i := 1; i < m+1 && s1[i-1] == s3[i-1]; i++ {
        dp[i][0] = true
    }
    // 初始化第一行，当s3前j个字符都由s2组成的情况
    for j := 1; j < n+1 && s2[j-1] == s3[j-1]; j++ {
        dp[0][j] = true
    }
    // 从第一个字符开始遍历，逐个字符比较计算
    for i := 1; i < m+1; i++ {
        for j := 1; j < n+1; j++ {
            k := i + j
            c1 := s1[i-1]
            c2 := s2[j-1]
            c3 := s3[k-1]
            // 有两种情况算匹配
            // 第一种，前一行匹配且s1当前字符与s3当前字符相同
            // 第二种，前一列匹配且s2当前字符与s3当前字符相同
            dp[i][j] = dp[i-1][j] && c1 == c3 || dp[i][j-1] && c2 == c3
        }
    }

    return dp[m][n]
}
