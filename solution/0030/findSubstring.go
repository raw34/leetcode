package _0030

func findSubstring(s string, words []string) []int {
    m := len(words)
    w := len(words[0])
    valid := func(window string) bool {
        hash := map[string]int{}
        for _, word := range words {
            hash[word] += 1
        }
        count, l, r := 0, 0, w
        for r <= len(window) {
            str := window[l:r]
            if _, ok := hash[str]; ok {
                hash[str]--
                count++
            }
            l += w
            r += w
        }
        for _, v := range hash {
            if v != 0 {
                return false
            }
        }

        return count == m
    }
    n := len(s)
    res := make([]int, 0)
    // 利用滑动窗口思路解决问题，窗口宽度固定为w*m
    for l, r := 0, w*m; r <= n; r++ {
        if valid(s[l:r]) {
            res = append(res, l)
        }
        l++
    }

    return res
}
