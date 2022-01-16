package runtime

import (
    "container/heap"
    "sort"
)

type MinHeap struct {
    sort.IntSlice
}

func NewMinHeap(nums []int) *MinHeap {
    mh := &MinHeap{}
    for _, num := range nums {
        heap.Push(mh, num)
    }
    return mh
}

func (this *MinHeap) Less(i, j int) bool {
    return this.IntSlice[i] < this.IntSlice[j]
}

func (this *MinHeap) Push(val interface{}) {
    this.IntSlice = append(this.IntSlice, val.(int))
}

func (this *MinHeap) Pop() interface{} {
    n := this.IntSlice.Len()
    val := this.IntSlice[n-1]
    this.IntSlice = this.IntSlice[:n-1]
    return val
}
