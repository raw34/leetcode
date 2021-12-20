package _0347

import "container/heap"

func topKFrequent(nums []int, k int) []int {
    // 利用哈希表对数字出现次数计数
    counts := map[int]int{}
    for _, num := range nums {
        counts[num]++
    }

    // 将哈希表构建为最小堆
    hp := &MinHeap{}
    heap.Init(hp)
    for key, count := range counts {
        heap.Push(hp, [2]int{key, count})
        if hp.Len() > k {
            heap.Pop(hp)
        }
    }

    // 逐个取出最小堆中剩余元素
    res := make([]int, hp.Len())
    for hp.Len() > 0 {
        num := heap.Pop(hp).([2]int)[0]
        res[hp.Len()] = num
    }

    return res
}

type MinHeap [][2]int

func (h MinHeap) Len() int {
    return len(h)
}
func (h MinHeap) Less(i, j int) bool {
    return h[i][1] < h[j][1]
}

func (h MinHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.([2]int))
}

func (h *MinHeap) Pop() interface{} {
    queue := *h
    x := queue[h.Len()-1]
    *h = queue[:h.Len()-1]
    return x
}
