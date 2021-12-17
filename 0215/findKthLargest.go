package _0215

import "github.com/raw34/leetcode/runtime"

func findKthLargest(nums []int, k int) int {
    queue := runtime.NewMaxPriorityQueue(nums)
    for i := 0; i < k-1; i++ {
        queue.Pop()
    }

    return queue.Pop()
}
