package _0703

import (
    "container/heap"
    "sort"
)

type KthLargest struct {
    k int
    sort.IntSlice
}

func Constructor(k int, nums []int) KthLargest {
    kl := KthLargest{k: k}
    for i := 0; i < len(nums); i++ {
        kl.Add(nums[i])
    }
    return kl
}

func (this *KthLargest) Add(val int) int {
    heap.Push(this, val)
    if this.Len() > this.k {
        heap.Pop(this)
    }
    return this.IntSlice[0]
}

func (this *KthLargest) Push(val interface{}) {
    this.IntSlice = append(this.IntSlice, val.(int))
}

func (this *KthLargest) Pop() interface{} {
    val := this.IntSlice[this.IntSlice.Len()-1]
    this.IntSlice = this.IntSlice[:this.IntSlice.Len()-1]
    return val
}
