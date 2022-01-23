package _0005

func longestPalindrome(s string) string {
    n := len(s)
    if n <= 1 {
        return s
    }
    valid := func(str string) bool {
        m := len(str)
        for i := 0; i < m; i++ {
            if str[i] != str[m-i-1] {
                return false
            }
        }
        return true
    }

    res := ""
    for i := 0; i < n; i++ {
        for j := i + 1; j < n+1; j++ {
            curr := s[i:j]
            if valid(curr) && len(curr) > len(res) {
                res = curr
            }
        }
    }

    return res
}

func longestPalindrome2(s string) string {
    n := len(s)
    if n <= 1 {
        return s
    }

    extend := func(s string, left, right int) int {
        for left >= 0 && right < len(s) && s[left] == s[right] {
            left--
            right++
        }
        return right - left - 1
    }
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }
    start, end := 0, 0
    for i := 0; i < n; i++ {
        len1 := extend(s, i, i)
        len2 := extend(s, i, i+1)
        maxLen := max(len1, len2)
        if maxLen > end-start {
            start = i - (maxLen-1)/2
            end = i + maxLen/2
        }
    }

    return s[start : end+1]
}
