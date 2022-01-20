package _0862

import "math"

func shortestSubarray(nums []int, k int) int {
    min := func(a, b int) int {
        if a > b {
            return b
        }
        return a
    }

    // 计算前缀和
    n := len(nums)
    preSum := make([]int, n+1)
    for i := 1; i < n+1; i++ {
        preSum[i] = preSum[i-1] + nums[i-1]
    }

    // 利用单调队列求解
    res := math.MaxInt32
    queue := make([]int, 0)
    for i := 0; i < n+1; i++ {
        curr := preSum[i]
        // 排除不可能情况，更新队列
        // 因为k为非负数，当前位置和前一个位置去区间和小于0时，必定找不到符合条件区间
        for len(queue) > 0 && curr <= preSum[queue[len(queue)-1]] {
            queue = queue[:len(queue)-1]
        }
        // 找到符合条件区间，更新结果并更新队列
        for len(queue) > 0 && curr-preSum[queue[0]] >= k {
            res = min(res, i-queue[0])
            queue = queue[1:]
        }
        queue = append(queue, i)
    }

    if res == math.MaxInt32 {
        res = -1
    }

    return res
}
