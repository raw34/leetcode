package _0739

func dailyTemperatures(temperatures []int) []int {
    n := len(temperatures)
    res := make([]int, n)
    stack := make([]int, 0)
    // 单调栈模板代码
    for i := 0; i < n; i++ {
        // 当栈不为空且当前元素大于栈顶元素时，计算当前元素与栈内元素的位移
        for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
            curr := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            res[curr] = i - curr
        }
        // 当栈空或当前元素小于等于栈顶元素时，入栈当前元素
        stack = append(stack, i)
    }

    return res
}
