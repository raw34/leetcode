package _0020

func isValid(s string) bool {
    n := len(s)
    hash := map[byte]byte{
        ')': '(',
        ']': '[',
        '}': '{',
    }
    stack := make([]byte, 0)
    for i := 0; i < n; i++ {
        curr := s[i]
        prev, ok := hash[curr]
        if len(stack) > 0 && ok && prev == stack[len(stack)-1] {
            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack, curr)
        }
    }

    return len(stack) == 0
}
