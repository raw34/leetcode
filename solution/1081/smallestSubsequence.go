package _1081

func smallestSubsequence(s string) string {
    // 核心思路，利用单调栈维护可保留字符串，可保留字符串条件为：
    // 1、不重复；2、不打乱顺序；3、字典序排列
    n := len(s)
    // 预处理，统计每个字符出现的次数
    counter := [256]int{}
    for i := 0; i < n; i++ {
        counter[s[i]]++
    }
    // 标记数组，用于判断栈中是否已存在字符
    inStack := [256]bool{}
    stack := make([]byte, 0)
    for i := 0; i < n; i++ {
        curr := s[i]
        // 当前字符计数减一
        counter[curr]--
        // 如果栈中已存在该字符，直接跳过
        if inStack[curr] {
            continue
        }
        // 如果当前字符小于栈顶字符，且栈顶字符后边还会出现，则弹出栈顶字符，并取消标记
        for len(stack) > 0 && curr < stack[len(stack)-1] && counter[stack[len(stack)-1]] > 0 {
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            inStack[top] = false
        }
        // 当前字符入栈，并标记
        stack = append(stack, curr)
        inStack[curr] = true
    }

    return string(stack)
}
