package _0037

func solveSudoku(board [][]byte) {
    m := len(board)
    n := len(board[0])
    // 行使用标记
    var rows [9][9]bool
    // 列使用标记
    var cols [9][9]bool
    // 九宫格使用标记
    var blocks [3][3][9]bool
    // 空白坐标数组
    spaces := make([][2]int, 0)

    // 获取所有空白坐标，初始化行、列、九宫格已使用标记
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            char := board[i][j]
            if char == '.' {
                spaces = append(spaces, [2]int{i, j})
            } else {
                char = char - '1'
                rows[i][char] = true
                cols[j][char] = true
                blocks[i/3][j/3][char] = true
            }
        }
    }

    var dfs func(index int) bool
    dfs = func(index int) bool {
        // 当空白坐标全部填充完毕，得到解
        if index == len(spaces) {
            return true
        }
        // 取出一个空白坐标
        i, j := spaces[index][0], spaces[index][1]
        // 从1到9尝试填充当前坐标
        for char := byte(0); char < byte(9); char++ {
            // 行、列、九宫格有一个被标记过则不合法
            if rows[i][char] || cols[j][char] || blocks[i/3][j/3][char] {
                continue
            }
            // 标记
            rows[i][char] = true
            cols[j][char] = true
            blocks[i/3][j/3][char] = true
            board[i][j] = char + '1'
            // 尝试下一个空白坐标，当下一个坐标得到解时直接退出
            if dfs(index + 1) {
                return true
            }
            // 回溯
            rows[i][char] = false
            cols[j][char] = false
            blocks[i/3][j/3][char] = false
        }

        return false
    }
    dfs(0)
}
