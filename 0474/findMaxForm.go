package _0474

import "strings"

func findMaxForm(strs []string, m int, n int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }

    l := len(strs)
    // dp[i][j][k]表示在前 i 个字符串中，使用 j 个 0 和 k 个 1 的情况下最多可以得到的字符串数量
    // 假设数组 strs 的长度为 l，则最终答案为 dp[l][m][n]
    dp := make([][][]int, l+1)
    for i := 0; i < l+1; i++ {
        dp[i] = make([][]int, m+1)
        for j := 0; j < m+1; j++ {
            dp[i][j] = make([]int, n+1)
        }
    }

    for i := 1; i < l+1; i++ {
        // 计算当前字符串0和1的个数
        str := strs[i-1]
        zeros := strings.Count(str, "0")
        ones := len(str) - zeros
        for j := 0; j < m+1; j++ {
            for k := 0; k < n+1; k++ {
                // 默认值为不选当前字符串的情况
                dp[i][j][k] = dp[i-1][j][k]
                if j >= zeros && k >= ones {
                    // 如果选择当前字符串，获取不选当前字符串且扣除0和1的空间后的前缀值+1，与默认值比较大小
                    dp[i][j][k] = max(dp[i][j][k], dp[i-1][j-zeros][k-ones]+1)
                }
            }
        }
    }

    return dp[l][m][n]
}

func findMaxForm2(strs []string, m int, n int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }

    l := len(strs)
    // dp[i][j]表示使用 j 个 0 和 k 个 1 的情况下最多可以得到的字符串数量，最终答案为 dp[m][n]
    dp := make([][]int, m+1)
    for i := 0; i < m+1; i++ {
        dp[i] = make([]int, n+1)
    }

    for i := 1; i < l+1; i++ {
        // 计算当前字符串0和1的个数
        str := strs[i-1]
        zeros := strings.Count(str, "0")
        ones := len(str) - zeros
        for j := m; j >= zeros; j-- {
            for k := n; k >= ones; k-- {
                // 默认值为不选当前字符串的情况
                dp[j][k] = dp[j][k]
                if j >= zeros && k >= ones {
                    // 如果选择当前字符串，获取不选当前字符串且扣除0和1的空间后的前缀值+1，与默认值比较大小
                    dp[j][k] = max(dp[j][k], dp[j-zeros][k-ones]+1)
                }
            }
        }
    }

    return dp[m][n]
}
