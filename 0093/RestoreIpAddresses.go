package _0093

import (
    "strconv"
    "strings"
)

func restoreIpAddresses(s string) []string {
    res := make([]string, 0)
    n := len(s)

    var backtrack func(subRes []string, start int)
    backtrack = func(subRes []string, start int) {
        if len(subRes) == 4 && start == n {
            res = append(res, strings.Join(subRes, "."))
            return
        }

        for length := 1; length < 4; length++ {
            end := start + length
            // 输入字符串不够长
            if end-1 >= n {
                return
            }
            // 当前段前缀为0
            if length != 1 && s[start] == '0' {
                return
            }

            str := s[start:end]
            // 当前段大于255
            num, _ := strconv.Atoi(str)
            if num > 255 {
                return
            }

            subRes = append(subRes, str)
            backtrack(subRes, end)
            subRes = subRes[0 : len(subRes)-1]
        }
    }
    backtrack([]string{}, 0)

    return res
}
