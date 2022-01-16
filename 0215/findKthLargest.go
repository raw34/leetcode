package _0215

import (
    "container/heap"
    "sort"
)

func findKthLargest(nums []int, k int) int {
    mh := NewMaxHeap(nums)
    for i := 0; i < len(nums); i++ {
        if i < k-1 {
            heap.Pop(mh)
        }
    }
    return heap.Pop(mh).(int)
}

type MaxHeap struct {
    sort.IntSlice
}

func NewMaxHeap(nums []int) *MaxHeap {
    mh := &MaxHeap{}
    for _, num := range nums {
        heap.Push(mh, num)
    }

    return mh
}

func (this *MaxHeap) Less(i, j int) bool {
    return this.IntSlice[i] > this.IntSlice[j]
}

func (this *MaxHeap) Push(val interface{}) {
    this.IntSlice = append(this.IntSlice, val.(int))
}

func (this *MaxHeap) Pop() interface{} {
    n := this.IntSlice.Len()
    val := this.IntSlice[n-1]
    this.IntSlice = this.IntSlice[:n-1]
    return val
}
