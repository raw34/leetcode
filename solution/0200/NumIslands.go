package _0200

func numIslands(grid [][]byte) int {
    if grid == nil || len(grid) == 0 {
        return 0
    }

    result := 0
    nx := len(grid)
    ny := len(grid[0])
    for x := 0; x < nx; x++ {
        for y := 0; y < ny; y++ {
            // 以某个陆地起点开始，标记一座岛屿
            if string(grid[x][y]) == "1" {
                result++
                markIslandDFS(grid, x, y)
            }
        }
    }

    return result
}

func markIslandDFS(grid [][]byte, x, y int) {
    nx := len(grid)
    ny := len(grid[0])

    if x < 0 || y < 0 || x >= nx || y >= ny || string(grid[x][y]) == "0" {
        return
    }

    grid[x][y] = []byte("0")[0]
    markIslandDFS(grid, x-1, y)
    markIslandDFS(grid, x+1, y)
    markIslandDFS(grid, x, y-1)
    markIslandDFS(grid, x, y+1)
}
