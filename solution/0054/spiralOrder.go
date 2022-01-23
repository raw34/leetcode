package _0054

import (
    "math"
)

func spiralOrder(matrix [][]int) []int {
    m := len(matrix)
    n := len(matrix[0])
    dirs := map[string][]int{
        "left":   {0, -1},
        "right":  {0, 1},
        "top":    {-1, 0},
        "bottom": {1, 0},
    }
    nextDirs := map[string][][]int{
        "left":   {dirs["right"], dirs["bottom"]},
        "right":  {dirs["left"], dirs["top"]},
        "top":    {dirs["bottom"], dirs["left"]},
        "bottom": {dirs["top"], dirs["right"]},
    }
    getNextDir := func(prev []int, i, j int) []int {
        diff := []int{prev[0] - i, prev[1] - j}
        from := ""
        for dir, dxy := range dirs {
            if dxy[0] == diff[0] && dxy[1] == diff[1] {
                from = dir
            }
        }
        to := nextDirs[from]
        res := make([]int, 0)
        for k := 0; k < len(to); k++ {
            x := i + to[k][0]
            y := j + to[k][1]
            if x >= 0 && y >= 0 && x < m && y < n && matrix[x][y] != math.MinInt32 {
                res = []int{x, y}
                break
            }
        }

        return res
    }

    res := []int{matrix[0][0]}
    matrix[0][0] = math.MinInt32
    prev := []int{0, 0}
    i, j := 0, 1
    if m > 1 && n == 1 {
        i, j = 1, 0
    }
    for len(res) < m*n {
        res = append(res, matrix[i][j])
        matrix[i][j] = math.MinInt32
        next := getNextDir(prev, i, j)
        prev = []int{i, j}
        if len(next) > 0 {
            i, j = next[0], next[1]
        }
    }

    return res
}
