package _0218

import (
    "container/heap"
    "sort"
)

func getSkyline(buildings [][]int) [][]int {
    n := len(buildings)
    points := make([][]int, 0)
    for i := 0; i < n; i++ {
        points = append(points, []int{buildings[i][0], -buildings[i][2]})
        points = append(points, []int{buildings[i][1], buildings[i][2]})
    }
    sort.Slice(points, func(i, j int) bool {
        if points[i][0] != points[j][0] {
            return points[i][0] < points[j][0]
        }
        return points[i][1] < points[j][1]
    })

    res := make([][]int, 0)
    curr, prev := 0, 0
    queue := &maxHeap{}
    heap.Push(queue, curr)
    for _, dir := range points {
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
