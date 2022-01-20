package _0931

func minFallingPathSum(matrix [][]int) int {
    min := func(arr ...int) int {
        res := arr[0]
        for _, num := range arr {
            if num < res {
                res = num
            }
        }

        return res
    }
    n := len(matrix)
    // dp[i][j]表示当前位置的最小路径和
    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, n)
    }
    // 初始化第一行
    for j := 0; j < n; j++ {
        dp[0][j] = matrix[0][j]
    }
    // 从第二行开始遍历，每次取前一行符合条件的最小值和当前值相加
    for i := 1; i < n; i++ {
        for j := 0; j < n; j++ {
            if j == 0 {
                dp[i][j] = min(dp[i-1][j], dp[i-1][j+1]) + matrix[i][j]
            } else if j == n-1 {
                dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + matrix[i][j]
            } else {
                dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i-1][j+1]) + matrix[i][j]
            }
        }
    }

    // 解是最后一行最小值
    return min(dp[n-1]...)
}
