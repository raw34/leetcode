package _0118

func generate(numRows int) [][]int {
    res := make([][]int, numRows)
    for i := 0; i < numRows; i++ {
        res[i] = make([]int, i+1)
    }

    res[0][0] = 1
    for i := 1; i < numRows; i++ {
        res[i][0] = 1
        res[i][i] = 1
        for j := 1; j < i; j++ {
            res[i][j] = res[i-1][j-1] + res[i-1][j]
        }
    }

    return res
}
