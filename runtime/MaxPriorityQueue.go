package runtime

import "fmt"

type MaxPriorityQueue struct {
    queue []int
    size  int
}

func NewMaxPriorityQueue(nums []int) MaxPriorityQueue {
    queue := make([]int, len(nums)+1)
    queue[0] = -1
    maxPQ := MaxPriorityQueue{queue, 0}
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
        olderChild := this.left(k)
        rightChild := this.right(k)
        if rightChild <= this.size && this.less(olderChild, rightChild) {
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
