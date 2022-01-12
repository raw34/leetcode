package _0239

func maxSlidingWindow(nums []int, k int) []int {
    n := len(nums)
    // 利用单调队列将数组中k位置前单调递增的数字入队
    queue := make([]int, 0)
    for i := 0; i < k; i++ {
        for len(queue) > 0 && nums[i] >= nums[queue[len(queue)-1]] {
            queue = queue[:len(queue)-1]
        }
        queue = append(queue, i)
    }
    // 从k开始寻找每个滑动窗口的最大值
    res := []int{nums[queue[0]]}
    for i := k; i < n; i++ {
        for len(queue) > 0 && nums[i] >= nums[queue[len(queue)-1]] {
            queue = queue[:len(queue)-1]
        }
        queue = append(queue, i)
        for queue[0] <= i-k {
            queue = queue[1:]
        }
        res = append(res, nums[queue[0]])
    }

    return res
}
