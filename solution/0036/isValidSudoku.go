package _0036

func isValidSudoku(board [][]byte) bool {
    n := len(board)
    dx := [9][9]bool{}
    dy := [9][9]bool{}
    dc := [3][3][9]bool{}

    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            curr := board[i][j]
            if curr == '.' {
                continue
            }
            curr -= '1'
            if dx[i][curr] || dy[j][curr] || dc[i/3][j/3][curr] {
                return false
            }
            dx[i][curr] = true
            dy[j][curr] = true
            dc[i/3][j/3][curr] = true
        }
    }

    return true
}
