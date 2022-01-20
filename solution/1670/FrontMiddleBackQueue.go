package _1670

import "fmt"

type FrontMiddleBackQueue struct {
    size  int
    queue []int
}

func Constructor() FrontMiddleBackQueue {
    return FrontMiddleBackQueue{0, []int{}}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
    this.queue = append([]int{val}, this.queue...)
    this.size++
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
    mid := this.size / 2
    left := this.queue[:mid]
    right := this.queue[mid:this.size]
    tempLeft := make([]int, len(left))
    tempRight := make([]int, len(right))
    copy(tempLeft, left)
    copy(tempRight, right)
    this.queue = append(append(tempLeft, val), tempRight...)
    this.size++
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
    this.queue = append(this.queue, val)
    this.size++
}

func (this *FrontMiddleBackQueue) PopFront() int {
    if this.size == 0 {
        return -1
    }
    val := this.queue[0]
    this.queue = this.queue[1:]
    this.size--
    return val
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
    if this.size == 0 {
        return -1
    }
    mid := (this.size - 1) / 2
    val := this.queue[mid]
    left := this.queue[:mid]
    right := this.queue[mid+1 : this.size]
    this.queue = append(left, right...)
    this.size--
    return val
}

func (this *FrontMiddleBackQueue) PopBack() int {
    if this.size == 0 {
        return -1
    }
    val := this.queue[this.size-1]
    this.queue = this.queue[0 : this.size-1]
    this.size--
    return val
}

func (this *FrontMiddleBackQueue) Display() {
    fmt.Println(this.queue)
}
