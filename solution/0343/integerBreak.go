package _0343

func integerBreak(n int) int {
    max := func(arr ...int) int {
        res := arr[0]
        for _, num := range arr {
            if num > res {
                res = num
            }
        }

        return res
    }
    dp := make([]int, n+1)
    for i := 2; i < n+1; i++ {
        currMax := 0
        for j := 1; j < i; j++ {
            currMax = max(currMax, (i-j)*j, dp[i-j]*j)
        }
        dp[i] = currMax
    }

    return dp[n]
}
