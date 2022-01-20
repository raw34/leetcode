package _0079

func exist(board [][]byte, word string) bool {
    m := len(board)
    n := len(board[0])
    used := make([][]bool, m)
    for i := 0; i < m; i++ {
        used[i] = make([]bool, n)
    }
    var check func(row, col, i int) bool
    check = func(row, col, i int) bool {
        if i == len(word) {
            return true
        }

        if row < 0 || col < 0 || row >= m || col >= n || used[row][col] || board[row][col] != word[i] {
            return false
        }

        used[row][col] = true
        if check(row+1, col, i+1) || check(row-1, col, i+1) || check(row, col+1, i+1) || check(row, col-1, i+1) {
            return true
        }
        used[row][col] = false
        return false
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if check(i, j, 0) {
                return true
            }
        }
    }

    return false
}
