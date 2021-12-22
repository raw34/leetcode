package _0529

func updateBoard(board [][]byte, click []int) [][]byte {
    row, col := click[0], click[1]
    if board[row][col] == 'M' {
        board[row][col] = 'X'
        return board
    }

    m := len(board)
    n := len(board[0])
    dr := []int{-1, -1, -1, 0, 0, 1, 1, 1}
    dc := []int{-1, 0, 1, -1, 1, -1, 0, 1}

    countM := func(r, c int) int {
        count := 0
        for i := 0; i < 8; i++ {
            nr := r + dr[i]
            nc := c + dc[i]
            if nr < 0 || nc < 0 || nr >= m || nc >= n {
                continue
            }
            if board[nr][nc] == 'M' {
                count++
            }
        }
        return count
    }

    var dfs func(r, c int)
    dfs = func(r, c int) {
        // 挖到地雷边缘坐标，标记并返回
        count := countM(r, c)
        if count > 0 {
            // 需要将数字先转字符串，再转字符
            board[r][c] = byte(count + '0')
            return
        }

        // 标记安全坐标
        board[r][c] = 'B'
        // 标记相邻安全坐标，并递归
        for i := 0; i < 8; i++ {
            nr := r + dr[i]
            nc := c + dc[i]
            if nr < 0 || nc < 0 || nr >= m || nc >= n || board[nr][nc] != 'E' {
                continue
            }
            dfs(nr, nc)
        }
    }
    dfs(row, col)

    return board
}
