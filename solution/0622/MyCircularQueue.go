package _0622

import "fmt"

type MyCircularQueue struct {
    size    int
    counter int
    queue   []int
}

func Constructor(k int) MyCircularQueue {
    queue := make([]int, k)
    for i := 0; i < k; i++ {
        queue[i] = -1
    }

    return MyCircularQueue{k, 0, queue}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.counter == this.size {
        return false
    }
    for i := 1; i <= this.size-1; i++ {
        this.queue[i-1] = this.queue[i]
    }
    this.queue[this.size-1] = value
    this.counter++
    return true
}

func (this *MyCircularQueue) DeQueue() bool {
    if this.counter == 0 {
        return false
    }
    this.queue[this.size-this.counter] = -1
    this.counter--
    return true
}

func (this *MyCircularQueue) Front() int {
    if this.counter == 0 {
        return this.queue[0]
    }
    return this.queue[this.size-this.counter]
}

func (this *MyCircularQueue) Rear() int {
    return this.queue[this.size-1]
}

func (this *MyCircularQueue) IsEmpty() bool {
    return this.counter == 0
}

func (this *MyCircularQueue) IsFull() bool {
    return this.counter == this.size
}

func (this *MyCircularQueue) Display() {
    fmt.Println(this.queue)
}
