package _0036

func isValidSudoku(board [][]byte) bool {
    n := len(board)
    dx := [9][9]bool{}
    dy := [9][9]bool{}
    dc := [3][3][9]bool{}

    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            char := board[i][j]
            if char == '.' {
                continue
            }
            char -= '1'
            if dx[i][char] || dy[j][char] || dc[i/3][j/3][char] {
                return false
            }
            dx[i][char] = true
            dy[j][char] = true
            dc[i/3][j/3][char] = true
        }
    }

    return true
}
