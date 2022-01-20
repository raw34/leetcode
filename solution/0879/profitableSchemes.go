package _0879

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }
    res := 0
    m := len(group)
    mod := int(1e9 + 7)
    //dp[i][j][k] 表示在前 i 个工作中选择了 j 个员工，并且满足工作利润至少为 k 的情况下的盈利计划的总数目。
    dp := make([][][]int, m+1)
    for i := 0; i < m+1; i++ {
        dp[i] = make([][]int, n+1)
        for j := 0; j < n+1; j++ {
            dp[i][j] = make([]int, minProfit+1)
        }
    }
    dp[0][0][0] = 1
    for i := 1; i < m+1; i++ {
        members, earn := group[i-1], profit[i-1]
        for j := 0; j < n+1; j++ {
            for k := 0; k < minProfit+1; k++ {
                if j < members {
                    dp[i][j][k] = dp[i-1][j][k]
                } else {
                    dp[i][j][k] = (dp[i-1][j][k] + dp[i-1][j-members][max(k-earn, 0)]) % mod
                }
            }
        }
    }

    for j := 0; j < n+1; j++ {
        res = (res + dp[m][j][minProfit]) % mod
    }

    return res
}

func profitableSchemes2(n int, minProfit int, group []int, profit []int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }
    m := len(group)
    mod := int(1e9 + 7)
    dp := make([][]int, n+1)
    for j := 0; j < n+1; j++ {
        dp[j] = make([]int, minProfit+1)
        dp[j][0] = 1
    }
    for i := 1; i < m+1; i++ {
        members, earn := group[i-1], profit[i-1]
        for j := n; j >= members; j-- {
            for k := minProfit; k >= 0; k-- {
                dp[j][k] = (dp[j][k] + dp[j-members][max(k-earn, 0)]) % mod
            }
        }
    }

    return dp[n][minProfit]
}
