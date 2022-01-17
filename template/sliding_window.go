package template

// 滑动窗口，模板题 Leetcode 209. 长度最小的子数组

func sliding_window() {
    s := []int{2, 3, 1, 2, 4, 3}
    var valid func() bool
    window := make([]int, 0)
    left, right := 0, 0
    for right < len(s) {
        window = append(window, s[right])
        right++

        for valid() {
            window = window[1:]
            left++
        }
    }
}
