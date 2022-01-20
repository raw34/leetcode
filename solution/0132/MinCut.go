package _0132

import "math"

func minCut(s string) int {
    min := func(a, b int) int {
        if a < b {
            return a
        }
        return b
    }

    n := len(s)
    dp1 := make([][]bool, n)
    for i := 0; i < n; i++ {
        dp1[i] = make([]bool, n)
        for j := 0; j < n; j++ {
            dp1[i][j] = true
        }
    }

    for i := n; i >= 0; i-- {
        for j := i + 1; j < n; j++ {
            dp1[i][j] = s[i] == s[j] && dp1[i+1][j-1]
        }
    }

    dp2 := make([]int, n)
    for i := 0; i < n; i++ {
        if dp1[0][i] {
            continue
        }
        dp2[i] = math.MaxInt32
        for j := 0; j < i; j++ {
            if dp1[j+1][i] {
                dp2[i] = min(dp2[i], dp2[j]+1)
            }
        }
    }

    return dp2[n-1]
}
