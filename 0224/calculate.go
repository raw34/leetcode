package _0224

import (
    "strconv"
)

func calculate(s string) int {
    // 追加前导0，以方便后续统一处理逻辑
    if s[0] == '-' {
        s = "0" + s
    }

    n := len(s)
    stack := make([]string, 0)

    // 追加数字到栈顶，根据情况进行计算或直接入栈
    appendNum := func(right string) {
        length := len(stack)
        if length > 0 && (stack[length-1] == "+" || stack[length-1] == "-") && stack[length-2] != "(" {
            opt := stack[length-1]
            left := stack[length-2]
            stack = stack[:length-2]
            num1, _ := strconv.Atoi(left)
            num2, _ := strconv.Atoi(right)
            if opt == "-" {
                num2 = -num2
            }
            sum := strconv.Itoa(num1 + num2)
            stack = append(stack, sum)
        } else {
            stack = append(stack, right)
        }
    }

    for i := 0; i < n; i++ {
        length := len(stack)
        curr := string(s[i])
        if curr == " " {
            // 当前字符为空格，忽略
            continue
        } else if curr == "+" || curr == "-" || curr == "(" {
            // 当前字符为 +  -  ( 直接入栈
            stack = append(stack, curr)
        } else if curr >= "0" && curr <= "9" {
            // 当前字符为数字时
            right := curr
            // 先处理前一个字符为 - 的情况
            if length > 1 && stack[length-1] == "-" && stack[length-2] == "(" {
                right = "-" + right
                stack = stack[:length-1]
            }
            // 再处理当前数字为多位的情况
            for i+1 < n && s[i+1] >= '0' && s[i+1] <= '9' {
                right += string(s[i+1])
                i++
            }
            // 调用函数，入栈
            appendNum(right)
        } else if curr == ")" {
            // 当前字符为 ） 时
            // 先从栈顶拿到当前数字
            right := stack[length-1]
            stack = stack[:length-1]
            length = len(stack)
            // 先处理前一个字符为 - 的情况，再移除最近一个 (
            if length > 1 && stack[length-1] == "-" && stack[length-2] == "(" {
                right = "-" + right
                stack = stack[:length-2]
            } else {
                stack = stack[:length-1]
            }
            // 调用函数，入栈
            appendNum(right)
        }
    }

    res, _ := strconv.Atoi(stack[0])
    return res
}
