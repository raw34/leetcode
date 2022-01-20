package _0295

import (
    "container/heap"
    "sort"
)

type MinHeap struct {
    sort.IntSlice
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

type MaxHeap struct {
    MinHeap
}

func (this *MaxHeap) Less(i, j int) bool {
    return this.IntSlice[i] > this.IntSlice[j]
}

type MedianFinder struct {
    minQ *MaxHeap // 用最大堆维护比中位数小的数
    maxQ *MinHeap // 用最小堆维护比中位数大的数
}

func Constructor() MedianFinder {
    return MedianFinder{&MaxHeap{}, &MinHeap{}}
}

func (this *MedianFinder) AddNum(num int) {
    if this.minQ.Len() == 0 || num <= this.minQ.IntSlice[0] {
        heap.Push(this.minQ, num)
        if this.maxQ.Len()+1 < this.minQ.Len() {
            heap.Push(this.maxQ, heap.Pop(this.minQ).(int))
        }
    } else {
        heap.Push(this.maxQ, num)
        if this.maxQ.Len() > this.minQ.Len() {
            heap.Push(this.minQ, heap.Pop(this.maxQ).(int))
        }
    }
}

func (this *MedianFinder) FindMedian() float64 {
    if this.minQ.Len() > this.maxQ.Len() {
        return float64(this.minQ.IntSlice[0])
    }

    return float64(this.minQ.IntSlice[0]+this.maxQ.IntSlice[0]) / 2

}
