package _0784

import "strings"

func letterCasePermutation(s string) []string {
    res := make([]string, 0)

    n := len(s)
    var backtrack func(list []string, begin int)
    backtrack = func(list []string, begin int) {
        if len(list) == n {
            res = append(res, strings.Join(list, ""))
            return
        }

        char := s[begin]
        list = append(list, string(char))
        backtrack(list, begin+1)
        list = list[:len(list)-1]

        if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
            char = char ^ (1 << 5)
            list = append(list, string(char))
            backtrack(list, begin+1)
            list = list[:len(list)-1]
        }
    }
    backtrack([]string{}, 0)

    return res
}
