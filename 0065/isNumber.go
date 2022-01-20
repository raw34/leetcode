package _0065

func isNumber(s string) bool {
    isNum := func(curr byte) bool {
        return curr >= '0' && curr <= '9'
    }
    isInt := func(str string) bool {
        numCount := 0
        optCount := 0
        for i := 0; i < len(str); i++ {
            curr := str[i]
            if isNum(curr) {
                numCount++
            } else if curr == '+' || curr == '-' {
                optCount++
                if i > 0 || optCount > 1 {
                    return false
                }
            } else {
                return false
            }
        }

        if numCount == 0 {
            return false
        }

        return true
    }
    isFloat := func(str string) bool {
        numCount := 0
        optCount := 0
        dotCount := 0
        for i := 0; i < len(str); i++ {
            curr := str[i]
            if isNum(curr) {
                numCount++
            } else if curr == '+' || curr == '-' {
                optCount++
                if i > 0 || optCount > 1 {
                    return false
                }
            } else if curr == '.' {
                dotCount++
                if dotCount > 1 {
                    return false
                }
            } else {
                return false
            }
        }

        if dotCount == 0 || numCount == 0 {
            return false
        }

        return true
    }

    if isInt(s) || isFloat(s) {
        return true
    }

    n := len(s)
    for i := 0; i < n; i++ {
        curr := s[i]
        if i > 0 && i < n-1 && (curr == 'e' || curr == 'E') {
            left := s[:i]
            right := s[i+1:]
            if (isInt(left) || isFloat(left)) && isInt(right) {
                return true
            }
        }
    }

    return false
}
