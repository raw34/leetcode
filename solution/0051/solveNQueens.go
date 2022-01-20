package _0051

func solveNQueens(n int) [][]string {
    res := make([][]string, 0)

    // 皇后坐标生成棋盘方法
    queen2board := func(queens []int) []string {
        board := make([]string, 0)
        for i := 0; i < n; i++ {
            row := make([]byte, n)
            for j := 0; j < n; j++ {
                row[j] = '.'
            }
            row[queens[i]] = 'Q'
            board = append(board, string(row))
        }

        return board
    }

    // 回溯方法
    var backtrack func(queens []int, row int, cols, diags1, diags2 map[int]bool)
    backtrack = func(queens []int, row int, cols, diags1, diags2 map[int]bool) {
        // 遍历行等于目标行数时保存可行方案
        if row == n {
            board := queen2board(queens)
            res = append(res, board)
            return
        }

        // 从左到右尝试当前行可行的皇后列坐标
        for col := 0; col < n; col++ {
            diag1 := row - col // 对角线数组坐标
            diag2 := row + col // 副对角线数组坐标

            // 利用列坐标和两个对角线坐标数组判断可行性
            if cols[col] || diags1[diag1] || diags2[diag2] {
                continue
            }

            // 当坐标可行时，保存选择
            queens[row] = col
            cols[col] = true
            diags1[diag1] = true
            diags2[diag2] = true

            // 回溯
            backtrack(queens, row+1, cols, diags1, diags2)

            // 清除选择
            queens[row] = -1
            cols[col] = false
            diags1[diag1] = false
            diags2[diag2] = false
        }
    }

    queens := make([]int, n)
    for i := 0; i < n; i++ {
        queens[i] = -1
    }
    cols := map[int]bool{}
    diags1 := map[int]bool{}
    diags2 := map[int]bool{}
    backtrack(queens, 0, cols, diags1, diags2)

    return res
}
