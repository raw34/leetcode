package _1044

// TLE版本
func longestDupSubstring(s string) string {
    n := len(s)
    cnt := map[string]int{}
    maxLen := 0
    res := ""
    for i := 1; i <= n; i++ {
        for j := 0; j < i; j++ {
            curr := s[j:i]
            cnt[curr]++
            if cnt[curr] > 1 && len(curr) > maxLen {
                maxLen = len(curr)
                res = curr
            }
        }
    }

    return res
}
