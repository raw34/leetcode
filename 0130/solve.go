package _0130

func solve(board [][]byte) {
    m := len(board)
    n := len(board[0])

    var dfs func(x, y int)
    dfs = func(x, y int) {
        if x < 0 || y < 0 || x >= m || y >= n || board[x][y] != 'O' {
            return
        }
        board[x][y] = 'A'
        dfs(x-1, y)
        dfs(x+1, y)
        dfs(x, y-1)
        dfs(x, y+1)
    }

    // 标记上下边界的O
    for x := 0; x < m; x++ {
        dfs(x, 0)
        dfs(x, n-1)
    }

    // 标记左右边界的O
    for y := 1; y < n-1; y++ {
        dfs(0, y)
        dfs(m-1, y)
    }

    // 遍历矩阵，将标记过的坐标改回O，未标记过的坐标改为X
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == 'A' {
                board[i][j] = 'O'
            } else if board[i][j] == 'O' {
                board[i][j] = 'X'
            }
        }
    }
}
