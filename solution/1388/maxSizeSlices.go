package _1388

func maxSizeSlices(slices []int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }
    // 思路类似打家劫舍II问题
    robMax := func(slices []int) int {
        n := len(slices)
        choice := (n + 1) / 3
        dp := make([][]int, n+1)
        dp[0] = make([]int, choice+1)
        for i := 1; i <= n; i++ {
            dp[i] = make([]int, choice+1)
            for j := 1; j <= choice; j++ {
                if i == 1 {
                    dp[i][j] = max(dp[i-1][j], slices[i-1])
                } else {
                    dp[i][j] = max(dp[i-1][j], dp[i-2][j-1]+slices[i-1])
                }
            }
        }

        return dp[n][choice]
    }

    // 情况1：拿第一块，不拿最后一块
    max1 := robMax(slices[:len(slices)-1])
    // 情况1：不拿第一块，拿最后一块
    max2 := robMax(slices[1:])

    return max(max1, max2)
}
