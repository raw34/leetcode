package _003

func lengthOfLongestSubstring(s string) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }

    res := 0
    m := map[byte]int{}
    n := len(s)
    r := -1

    for l := 0; l < n; l++ {
        if l > 0 {
            delete(m, s[l-1])
        }

        for r+1 < n && m[s[r+1]] == 0 {
            m[s[r+1]]++
            r++
        }

        res = max(res, r-l+1)
    }

    return res
}
