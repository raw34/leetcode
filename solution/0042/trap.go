package _0042

func trap(height []int) int {
    min := func(a, b int) int {
        if a > b {
            return b
        }
        return a
    }
    stack := make([]int, 0)
    res := 0
    for i := 0; i < len(height); i++ {
        // 当前元素比栈顶元素高时，触发计算
        for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
            // 获取栈顶元素
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if len(stack) == 0 {
                break
            }
            // 获取前一个元素
            left := stack[len(stack)-1]
            // 计算当前矩形宽度
            currW := i - left - 1
            // 计算当前矩形高度，取左右柱高最小值
            currH := min(height[left], height[i]) - height[top]
            // 累加总面积
            res += currW * currH
        }
        // 当前元素入栈
        stack = append(stack, i)
    }

    return res
}
