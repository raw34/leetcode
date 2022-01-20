package _0402

import "strings"

func removeKdigits(num string, k int) string {
    n := len(num)
    remain := n - k
    // 利用单调栈保存单调递增的数字
    stack := make([]byte, 0)
    for i := 0; i < n; i++ {
        curr := num[i]
        // 当移除元素个数不够且当前元素小于栈顶元素时，栈顶元素出栈，移除元素计数更新
        for k > 0 && len(stack) > 0 && curr < stack[len(stack)-1] {
            stack = stack[:len(stack)-1]
            k--
        }
        // 当前元素入栈
        stack = append(stack, curr)
    }
    // 截取部分元素，避免移除元素数量不够的情况
    stack = stack[:remain]
    // 规避前导0的情况
    res := strings.TrimLeft(string(stack), "0")
    if res == "" {
        res = "0"
    }

    return res
}
