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
    dxy := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
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
            if x < 0 || y < 0 || x > m-1 || y > n-1 || matrix[i][j] >= matrix[x][y] {
                continue
            }
            memo[i][j] = max(memo[i][j], dfs(x, y, memo)+1)
        }
        return memo[i][j]
    }
    res := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            res = max(res, dfs(i, j, memo))
        }
    }

    return res
}
