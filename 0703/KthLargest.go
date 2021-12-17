package _0703

import "github.com/raw34/leetcode/runtime"

type KthLargest struct {
    k   int
    mpq *runtime.MaxPriorityQueue
}

func Constructor(k int, nums []int) KthLargest {
    return KthLargest{k, runtime.NewMaxPriorityQueue(nums)}
}

func (this *KthLargest) Add(val int) int {
    return val
}
