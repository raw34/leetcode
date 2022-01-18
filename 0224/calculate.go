package _0224

import (
    "strconv"
)

func calculate(s string) int {
    if s[0] == '-' {
        s = "0" + s
    }

    n := len(s)
    stack := make([]string, 0)

    updateStack := func(right string) {
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
            continue
        } else if curr == "+" || curr == "-" || curr == "(" {
            stack = append(stack, curr)
        } else if curr >= "0" && curr <= "9" {
            right := curr
            if length > 1 && stack[length-1] == "-" && stack[length-2] == "(" {
                right = "-" + right
                stack = stack[:length-1]
            }
            for i+1 < n && s[i+1] >= '0' && s[i+1] <= '9' {
                right += string(s[i+1])
                i++
            }
            updateStack(right)
        } else if curr == ")" {
            right := stack[length-1]
            stack = stack[:length-1]
            length = len(stack)
            if length > 1 && stack[length-1] == "-" && stack[length-2] == "(" {
                right = "-" + right
                stack = stack[:length-2]
            } else {
                stack = stack[:length-1]
            }
            updateStack(right)
        }
    }

    res, _ := strconv.Atoi(stack[0])
    return res
}
