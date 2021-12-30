package _0072

func minDistance(word1 string, word2 string) int {
    min := func(arr ...int) int {
        res := arr[0]
        for _, num := range arr {
            if num < res {
                res = num
            }
        }
        return res
    }

    n1 := len(word1)
    n2 := len(word2)
    // dp[i][j] 表示A串i位置字符到B串j位置字符的编辑距离
    dp := make([][]int, n1+1)
    for i := 0; i <= n1; i++ {
        dp[i] = make([]int, n2+1)
    }

    // 初始化A串每个位置到B串第一个字符的编辑距离
    for i := 1; i <= n1; i++ {
        dp[i][0] = dp[i-1][0] + 1
    }

    // 初始化B串每个位置到A串第一个字符的编辑距离
    for j := 1; j <= n2; j++ {
        dp[0][j] = dp[0][j-1] + 1
    }

    // 从A串第二个字符和B串第二个字符开始遍历
    for i := 1; i <= n1; i++ {
        for j := 1; j <= n2; j++ {
            if word1[i-1] == word2[j-1] {
                // 当两串当前字符相同时不需要增加编辑举例
                dp[i][j] = dp[i-1][j-1]
            } else {
                // 当两串当前字符不同时需要增加编辑距离，最小的临近情况 + 1
                dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
            }
        }
    }

    return dp[n1][n2]
}
