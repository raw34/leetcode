package _0218

import (
    "container/heap"
    "sort"
)

func getSkyline(buildings [][]int) [][]int {
    n := len(buildings)
    dirs := make([][]int, 0)
    for i := 0; i < n; i++ {
        dirs = append(dirs, []int{buildings[i][0], -buildings[i][2]})
        dirs = append(dirs, []int{buildings[i][1], buildings[i][2]})
    }
    sort.Slice(dirs, func(i, j int) bool {
        if dirs[i][0] != dirs[j][0] {
            return dirs[i][0] < dirs[j][0]
        }
        return dirs[i][1] < dirs[j][1]
    })

    res := make([][]int, 0)
    queue := &maxHeap{}
    heap.Push(queue, 0)
    curr, prev := 0, 0
    for _, dir := range dirs {
        x, y := dir[0], dir[1]
        if y < 0 {
            heap.Push(queue, -y)
        } else {
            heap.Remove(queue, queue.Find(y))
        }

        if queue.Len() > 0 {
            curr = heap.Pop(queue).(int)
            heap.Push(queue, curr)
        }
        if curr != prev {
            res = append(res, []int{x, curr})
            prev = curr
        }
    }

    return res
}

type maxHeap struct {
    sort.IntSlice
}

func (this *maxHeap) Less(i, j int) bool {
    return this.IntSlice[i] > this.IntSlice[j]
}

func (this *maxHeap) Push(val interface{}) {
    this.IntSlice = append(this.IntSlice, val.(int))
}

func (this *maxHeap) Pop() interface{} {
    n := this.IntSlice.Len()
    val := this.IntSlice[n-1]
    this.IntSlice = this.IntSlice[:n-1]
    return val
}

func (this *maxHeap) Find(val interface{}) int {
    idx := -1
    for i := 0; i < this.IntSlice.Len(); i++ {
        if this.IntSlice[i] == val.(int) {
            idx = i
            break
        }
    }
    return idx
}
