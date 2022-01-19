package _0329

func longestIncreasingPath(matrix [][]int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }
    m := len(matrix)
    n := len(matrix[0])
    // 四个可移动方向的坐标
    dxy := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
    // 备忘录，表示从matrix[i][j]为起点的路径长度，不带备忘录会LTE
    memo := make([][]int, m)
    for i := 0; i < m; i++ {
        memo[i] = make([]int, n)
    }
    var dfs func(i, j int, memo [][]int) int
    dfs = func(i, j int, memo [][]int) int {
        if memo[i][j] != 0 {
            return memo[i][j]
        }
        memo[i][j]++
        for _, xy := range dxy {
            x := i + xy[0]
            y := j + xy[1]
            // 越界或不符合条件时跳过
            if x < 0 || y < 0 || x > m-1 || y > n-1 || matrix[i][j] >= matrix[x][y] {
                continue
            }
            memo[i][j] = max(memo[i][j], dfs(x, y, memo)+1)
        }
        return memo[i][j]
    }

    // 以矩阵的所有坐标做起点开始遍历
    res := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            res = max(res, dfs(i, j, memo))
        }
    }

    return res
}
