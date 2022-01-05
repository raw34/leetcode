package _0264

func nthUglyNumber(n int) int {
    min := func(arr ...int) int {
        res := arr[0]
        for _, num := range arr {
            if num < res {
                res = num
            }
        }
        return res
    }
    x, y, z := 0, 0, 0
    dp := make([]int, n)
    dp[0] = 1
    for i := 1; i < n; i++ {
        num2, num3, num5 := dp[x]*2, dp[y]*3, dp[z]*5
        dp[i] = min(num2, num3, num5)
        if dp[i] == num2 {
            x++
        }
        if dp[i] == num3 {
            y++
        }
        if dp[i] == num5 {
            z++
        }
    }

    return dp[n-1]
}
