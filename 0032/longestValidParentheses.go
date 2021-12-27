package _0032

func longestValidParentheses(s string) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }

    n := len(s)
    res := 0
    dp := make([]int, n)
    for i := 1; i < n; i++ {
        // 只有出现右括号才能组成完整括号，左括号直接跳过
        if s[i] == '(' {
            continue
        }
        // 获取前一个位置计算好的完整括号长度
        prevMax := dp[i-1]
        // 找到前驱的完整括号之前的索引位置
        prevIndex := i - prevMax
        if s[i-1] == '(' {
            // 如果前一个位置为左括号，给当前位置赋初始长度
            dp[i] = 2
            if i-2 >= 0 {
                // 如果当前一对括号之前索引存在，累加长度值
                dp[i] = dp[i] + dp[i-2]
            }
        } else if prevIndex-1 >= 0 && s[prevIndex-1] == '(' {
            // 如果前驱索引存在，且该位置为左括号，给当前位置赋值，在前驱完整括号长度上加2
            dp[i] = prevMax + 2
            if prevIndex-2 >= 0 {
                // 如果前驱索引之前索引存在，继续累加长度值
                dp[i] = dp[i] + dp[prevIndex-2]
            }
        }
        res = max(res, dp[i])
    }

    return res
}
