package _542

func updateMatrix(mat [][]int) [][]int {
	ni := len(mat)
	nj := len(mat[0])

	// 找到所有0的位置
	queue := make([][]int, 0)
	for i := 0; i < ni; i++ {
		for j := 0; j < nj; j++ {
			if mat[i][j] == 0 {
				queue = append(queue, []int{i, j})
			} else {
				mat[i][j] = -1
			}
		}
	}

	// 计算所有非0位置的位移
	direction := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]

		for _, dir := range direction {
			i := cell[0] + dir[0]
			j := cell[1] + dir[1]

			if i >= 0 && i < ni && j >= 0 && j < nj && mat[i][j] == -1 {
				mat[i][j] = mat[cell[0]][cell[1]] + 1
				queue = append(queue, []int{i, j})
			}
		}
	}

	return mat
}