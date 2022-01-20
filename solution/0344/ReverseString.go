package _0344

func reverseString(s []byte) {
    var reverse func(s []byte, left, right int)
    reverse = func(s []byte, left, right int) {
        if left >= right {
            return
        }

        temp := s[left]
        s[left] = s[right]
        s[right] = temp

        reverse(s, left+1, right-1)
    }

    reverse(s, 0, len(s)-1)
}
