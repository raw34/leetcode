package _0131

func partition(s string) [][]string {
    n := len(s)
    dp := make([][]bool, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]bool, n)
        for j := 0; j < n; j++ {
            dp[i][j] = true
        }
    }

    for i := n; i >= 0; i-- {
        for j := i + 1; j < n; j++ {
            dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
        }
    }

    res := make([][]string, 0)
    var backtrack func(list []string, i int)
    backtrack = func(list []string, i int) {
        if i == n {
            temp := make([]string, len(list))
            copy(temp, list)
            res = append(res, temp)
        }

        for j := i; j < n; j++ {
            if dp[i][j] {
                list = append(list, s[i:j+1])
                backtrack(list, j+1)
                list = list[0 : len(list)-1]
            }
        }
    }
    backtrack([]string{}, 0)

    return res
}
