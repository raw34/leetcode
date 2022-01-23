package _0007

import "math"

func reverse(x int) int {
    y := 0
    for x != 0 {
        if y < math.MinInt32/10 || y > math.MaxInt32/10 {
            return 0
        }
        digit := x % 10
        x /= 10
        y = y*10 + digit
    }

    return y
}
