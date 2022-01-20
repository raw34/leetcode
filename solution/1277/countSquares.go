package _1277

func countSquares(matrix [][]int) int {
    min := func(arr ...int) int {
        res := arr[0]
        for _, num := range arr {
            if num < res {
                res = num
            }
        }

        return res
    }
    m := len(matrix)
    n := len(matrix[0])
    res := 0
    dp := make([][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] != 1 {
                continue
            }
            if i == 0 || j == 0 {
                dp[i][j] = 1
            } else {
                dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
            }
            res += dp[i][j]
        }
    }

    return res
}
