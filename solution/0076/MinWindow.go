package _0076

import "math"

func minWindow(s string, t string) string {
    std := map[byte]int{}
    cnt := map[byte]int{}
    for i := 0; i < len(t); i++ {
        std[t[i]]++
    }

    match := func() bool {
        for k, v := range std {
            if cnt[k] < v {
                return false
            }
        }
        return true
    }

    ns := len(s)
    resL := -1
    resLen := math.MaxInt32
    for l, r := 0, 0; r < ns; r++ {
        if r < ns {
            cnt[s[r]]++
        }
        for match() && l <= r {
            currLen := r - l + 1
            if currLen < resLen {
                resL = l
                resLen = currLen
            }
            cnt[s[l]]--
            l++
        }
    }

    if resL == -1 {
        return ""
    }

    return s[resL : resL+resLen]
}
