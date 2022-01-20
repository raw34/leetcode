package _0650

import "math"

func minSteps(n int) int {
    min := func(arr ...int) int {
        res := arr[0]
        for _, num := range arr {
            if num < res {
                res = num
            }
        }

        return res
    }
    dp := make([]int, n+1)
    for i := 2; i < n+1; i++ {
        dp[i] = math.MaxInt32
        for j := 1; j*j < i+1; j++ {
            if i%j == 0 {
                dp[i] = min(dp[i], dp[j]+i/j, dp[i/j]+j)
            }
        }
    }

    return dp[n]
}

func minSteps2(n int) int {
    res := 0
    for i := 2; i < n+1; i++ {
        for n%i == 0 {
            n /= i
            res += i
        }
    }
    if n > 1 {
        res += n
    }

    return res
}
