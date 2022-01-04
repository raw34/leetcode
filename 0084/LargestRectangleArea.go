package _0084

func largestRectangleArea(heights []int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }

    area := 0
    // 数组前后各追加一个0高度，以方便计算边界
    heights = append([]int{0}, append(heights, 0)...)
    // 栈初始放一个0高度，避免非空判断
    stack := []int{0}
    // 从头遍历数组，计算数组中每个高度所能构成矩形的最大值
    for i := 0; i < len(heights); i++ {
        // 遇到当前元素高度小于栈顶元素高度，可以计算栈中高度的矩形面积
        for heights[i] < heights[stack[len(stack)-1]] {
            // 获取栈顶元素高度
            curr := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            currH := heights[curr]

            // 获取栈顶左侧元素索引
            left := stack[len(stack)-1]
            // 右侧元素索引为当前位置
            right := i
            // 计算栈顶元素宽度
            currW := right - left - 1
            // 计算栈顶元素宽高所对应的面积
            area = max(area, currW*currH)
        }

        // 如果当前元素高度大于等于栈顶元素高度，入栈
        stack = append(stack, i)
    }

    return area
}
