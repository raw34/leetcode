package _003

func lengthOfLongestSubstring(s string) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }

    res := 0
    cnt := map[byte]int{}
    n := len(s)

    for l, r := 0, 0; l < n; l++ {
        if l > 0 {
            delete(cnt, s[l-1])
        }

        for r < n && cnt[s[r]] == 0 {
            cnt[s[r]]++
            r++
        }

        res = max(res, r-l)
    }

    return res
}
