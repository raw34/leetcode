package _0218

import (
    "container/heap"
    "sort"
)

func getSkyline(buildings [][]int) [][]int {
    n := len(buildings)
    boundaries := make([]int, 0)
    for i := 0; i < n; i++ {
        boundaries = append(boundaries, buildings[i][0], buildings[i][1])
    }
    sort.Ints(boundaries)
    res := make([][]int, 0)
    queue := maxHeap{}
    idx := 0
    for _, boundary := range boundaries {
        for idx < n && buildings[idx][0] <= boundary {
            heap.Push(&queue, [2]int{buildings[idx][1], buildings[idx][2]})
            idx++
        }
        for queue.Len() > 0 && queue[0][0] <= boundary {
            heap.Pop(&queue)
        }
        maxHeight := 0
        if queue.Len() > 0 {
            maxHeight = queue[0][1]
        }
        if len(res) == 0 || maxHeight != res[len(res)-1][1] {
            res = append(res, []int{boundary, maxHeight})
        }
    }

    return res
}

type maxHeap [][2]int

func (h maxHeap) Len() int {
    return len(h)
}

func (h maxHeap) Less(i, j int) bool {
    return h[i][1] > h[j][1]
}

func (h maxHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Push(val interface{}) {
    *h = append(*h, val.([2]int))
}

func (h *maxHeap) Pop() interface{} {
    queue := *h
    n := queue.Len()
    val := queue[n-1]
    *h = queue[:n-1]
    return val
}
