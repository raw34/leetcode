package _0065

func isNumber(s string) bool {
    // 判断当前字符是否为单数字
    isNum := func(curr byte) bool {
        return curr >= '0' && curr <= '9'
    }
    // 判断字符串是否为纯整数或纯小数
    isIntOrFloat := func(str string, optMax, dotMax int) bool {
        numCount := 0
        optCount := 0
        dotCount := 0
        for i := 0; i < len(str); i++ {
            curr := str[i]
            if isNum(curr) {
                numCount++
            } else if curr == '+' || curr == '-' {
                optCount++
                if i > 0 || optCount > optMax {
                    return false
                }
            } else if curr == '.' {
                dotCount++
                if dotCount > optMax {
                    return false
                }
            } else {
                return false
            }
        }

        if numCount == 0 || dotCount != dotMax {
            return false
        }

        return true
    }
    // 判断当前字符串是否为纯整数
    isInt := func(str string) bool {
        return isIntOrFloat(str, 1, 0)
    }
    // 判断当前字符串是否为纯小数
    isFloat := func(str string) bool {
        return isIntOrFloat(str, 1, 1)
    }

    // 当字符串为纯整数或纯小数，直接返回
    if isInt(s) || isFloat(s) {
        return true
    }

    // 处理字符串中带e的情况
    n := len(s)
    for i := 0; i < n; i++ {
        curr := s[i]
        // 当前字符串为e，且e左右都为合法数字时返回
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
