package _0363

import "math"

func maxSumSubmatrix(matrix [][]int, k int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }

    m := len(matrix)
    n := len(matrix[0])
    // 计算前缀和
    preSum := make([][]int, m+1)
    for i := 0; i < m+1; i++ {
        preSum[i] = make([]int, n+1)
    }
    for i := 1; i < m+1; i++ {
        for j := 1; j < n+1; j++ {
            preSum[i][j] = preSum[i-1][j] + preSum[i][j-1] - preSum[i-1][j-1] + matrix[i-1][j-1]
        }
    }

    // 暴力+前缀和求每个矩形和
    res := math.MinInt32
    for i := 1; i < m+1; i++ {
        for j := 1; j < n+1; j++ {
            for p := i; p < m+1; p++ {
                for q := j; q < n+1; q++ {
                    currSum := preSum[p][q] - preSum[i-1][q] - preSum[p][j-1] + preSum[i-1][j-1]
                    if currSum <= k {
                        res = max(res, currSum)
                    }
                }
            }
        }
    }

    return res
}
