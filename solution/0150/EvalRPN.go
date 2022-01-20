package _0150

import "strconv"

func evalRPN(tokens []string) int {
    stack := make([]int, 0)

    for _, token := range tokens {
        num, err := strconv.Atoi(token)
        if err == nil {
            stack = append(stack, num)
            continue
        }

        length := len(stack)
        left, right := stack[length-2], stack[length-1]
        stack = stack[0 : length-2]

        switch token {
        case "+":
            stack = append(stack, left+right)
        case "-":
            stack = append(stack, left-right)
        case "*":
            stack = append(stack, left*right)
        case "/":
            stack = append(stack, left/right)
        }
    }

    return stack[0]
}
