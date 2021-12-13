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
    size  int
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
    this.size++
    this.queue[this.size] = val
    this.swim(this.size)
}

func (this *MaxPriorityQueue) swim(k int) {
    for k > 1 && this.less(this.parent(k), k) {
        this.swap(this.parent(k), k)
        k = this.parent(k)
    }
}

func (this *MaxPriorityQueue) swap(i, j int) {
    this.queue[i], this.queue[j] = this.queue[j], this.queue[i]
}

func (this *MaxPriorityQueue) Pop() int {
    max := this.queue[1]
    this.swap(1, this.size)
    this.queue = this.queue[0:this.size]
    this.size--
    this.sink(1)
    return max
}

func (this *MaxPriorityQueue) sink(k int) {
    for this.left(k) <= this.size {
        leftK := this.left(k)
        rightK := this.right(k)
        if rightK <= this.size && this.less(leftK, rightK) {
            leftK = rightK
        }

        if this.less(leftK, k) {
            break
        }

        this.swap(k, leftK)
        k = leftK
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
