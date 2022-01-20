package _0340

func lengthOfLongestSubstringKDistinct(s string, k int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }
    n := len(s)
    hash := map[byte]int{}
    res := 0
    for l, r := 0, 0; r < n; r++ {
        hash[s[r]]++
        count := len(hash)
        if count > k {
            hash[s[l]]--
            if hash[s[l]] == 0 {
                delete(hash, s[l])
            }
            l++
        }
        res = max(res, r-l+1)
    }

    return res
}
