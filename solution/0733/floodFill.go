package _0733

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    m := len(image)
    n := len(image[0])
    oldColor := image[sr][sc]
    dx := []int{1, 0, 0, -1}
    dy := []int{0, 1, -1, 0}

    var backtrack func(sr, sc int)
    backtrack = func(sr, sc int) {
        if image[sr][sc] != oldColor {
            return
        }

        image[sr][sc] = newColor
        for i := 0; i < 4; i++ {
            x := sr + dx[i]
            y := sc + dy[i]

            if x >= 0 && x < m && y >= 0 && y < n {
                backtrack(x, y)
            }
        }
    }
    backtrack(sr, sc)

    return image
}
