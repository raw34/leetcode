package _0014

func longestCommonPrefix(strs []string) string {
    m := len(strs)
    n := len(strs[0])
    prefix := ""
    for i := 0; i < n; i++ {
        c := strs[0][i]
        for j := 1; j < m; j++ {
            curr := strs[j]
            if i > len(curr)-1 || strs[j][i] != c {
                return prefix
            }
        }
        prefix += string(c)
    }

    return prefix
}
