package _0066

func plusOne(digits []int) []int {
    i := len(digits)
    carry := 1
    for carry > 0 && i > 0 {
        i--
        sum := digits[i] + 1
        digits[i] = sum % 10
        carry = sum / 10
    }
    if carry > 0 {
        digits = append([]int{carry}, digits...)
    }

    return digits
}
