package _0011

func maxArea(height []int) int {
    min := func(a, b int) int {
        if a > b {
            return b
        }
        return a
    }
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }
    res := 0
    l := 0
    r := len(height) - 1
    for l < r {
        area := min(height[l], height[r]) * (r - l)
        res = max(res, area)
        if height[l] <= height[r] {
            l++
        } else {
            r--
        }
    }

    return res
}
