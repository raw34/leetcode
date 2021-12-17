package runtime

import "fmt"

type MaxPriorityQueue struct {
    queue []int
}

func NewMaxPriorityQueue(nums []int) MaxPriorityQueue {
    queue := make([]int, 1)
    queue[0] = -1
    maxPQ := MaxPriorityQueue{queue}
    for i := 0; i < len(nums); i++ {
        maxPQ.Push(nums[i])
    }
    return maxPQ
}

func (this *MaxPriorityQueue) Push(val int) {
    this.queue = append(this.queue, val)
    this.swim(len(this.queue) - 1)
}

func (this *MaxPriorityQueue) Pop() int {
    max := this.queue[1]
    this.swap(1, len(this.queue)-1)
    this.queue = this.queue[:len(this.queue)-1]
    this.sink(1)
    return max
}

func (this *MaxPriorityQueue) swim(k int) {
    for k > 1 && this.less(this.parent(k), k) {
        this.swap(this.parent(k), k)
        k = this.parent(k)
    }
}

func (this *MaxPriorityQueue) less(i, j int) bool {
    return this.queue[i] < this.queue[j]
}

func (this *MaxPriorityQueue) swap(i, j int) {
    this.queue[i], this.queue[j] = this.queue[j], this.queue[i]
}

func (this *MaxPriorityQueue) sink(k int) {
    for this.left(k) <= len(this.queue)-1 {
        olderChild := this.left(k)
        rightChild := this.right(k)
        if rightChild <= len(this.queue)-1 && this.less(olderChild, rightChild) {
            olderChild = rightChild
        }

        if this.less(olderChild, k) {
            break
        }

        this.swap(k, olderChild)
        k = olderChild
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
