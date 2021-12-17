package _0703

import "github.com/raw34/leetcode/runtime"

type KthLargest struct {
    k   int
    mpq runtime.MaxPriorityQueue
}

func Constructor(k int, nums []int) KthLargest {
    return KthLargest{k, runtime.NewMaxPriorityQueue(nums)}
}

func (this *KthLargest) Add(val int) int {
    this.mpq.Push(val)
    return this.Search(this.k)
}

func (this *KthLargest) Search(k int) int {
    queue := this.mpq
    for i := 0; i < k-1; i++ {
        queue.Pop()
    }

    return queue.Pop()
}
