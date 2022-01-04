package _0085

func maximalRectangle(matrix [][]byte) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }

    getMaxArea := func(heights []int) int {
        area := 0
        stack := []int{0}
        for i := 0; i < len(heights); i++ {
            for heights[i] < heights[stack[len(stack)-1]] {
                curr := stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                currH := heights[curr]

                left := stack[len(stack)-1]
                right := i
                currW := right - left - 1
                area = max(area, currW*currH)
            }
            stack = append(stack, i)
        }
        return area
    }

    m := len(matrix)
    n := len(matrix[0])
    res := 0
    // 逐行计算柱状图最大矩形面积
    heights := make([]int, n+2)
    for i := 0; i < m; i++ {
        // 将矩阵转换为柱状图，可以将问题转化为求柱状图最大矩形面积（No.84）
        for j := 0; j < n; j++ {
            if matrix[i][j] == '0' {
                heights[j+1] = 0
            } else {
                heights[j+1] += 1
            }
        }
        res = max(res, getMaxArea(heights))
    }

    return res
}
