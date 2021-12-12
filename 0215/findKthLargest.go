package _0215

import "fmt"

func findKthLargest(nums []int, k int) int {
    queue := Constructor(nums)
    for i := 0; i < k-1; i++ {
        queue.Pop()
    }

    return queue.Pop()
}

type MaxPriorityQueue struct {
    queue []int
    count int
}

func Constructor(nums []int) *MaxPriorityQueue {
    queue := make([]int, len(nums)+1)
    queue[0] = -1
    maxPQ := &MaxPriorityQueue{queue, 0}
    for i := 0; i < len(nums); i++ {
        maxPQ.Push(nums[i])
    }
    return maxPQ
}

func (this *MaxPriorityQueue) less(i, j int) bool {
    return this.queue[i] < this.queue[j]
}

func (this *MaxPriorityQueue) Push(val int) {
    this.count++
    this.queue[this.count] = val
    this.swim(this.count)
}

func (this *MaxPriorityQueue) swim(i int) {
    for i > 1 && this.less(this.parent(i), i) {
        this.swap(this.parent(i), i)
        i = this.parent(i)
    }
}

func (this *MaxPriorityQueue) swap(i, j int) {
    this.queue[i], this.queue[j] = this.queue[j], this.queue[i]
}

func (this *MaxPriorityQueue) Pop() int {
    max := this.queue[1]
    this.swap(1, this.count)
    this.queue = this.queue[0:this.count]
    this.count--
    this.sink(1)
    return max
}

func (this *MaxPriorityQueue) sink(i int) {
    for this.left(i) <= this.count {
        left := this.left(i)
        right := this.right(i)
        if right <= this.count && this.less(left, right) {
            left = right
        }

        if this.less(left, i) {
            break
        }

        this.swap(i, left)
        i = left
    }
}

func (this *MaxPriorityQueue) parent(k int) int {
    return k / 2
}

func (this *MaxPriorityQueue) left(k int) int {
    return k * 2
}

func (this *MaxPriorityQueue) right(k int) int {
    return k*2 + 1
}

func (this *MaxPriorityQueue) Display() {
    fmt.Println(this.queue)
}
