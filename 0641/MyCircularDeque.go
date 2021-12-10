package _0641

import "fmt"

type MyCircularDeque struct {
    size    int
    counter int
    queue   []int
}

func Constructor(k int) MyCircularDeque {
    queue := make([]int, 0)
    return MyCircularDeque{k, 0, queue}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
    if this.counter == this.size {
        return false
    }
    this.queue = append([]int{value}, this.queue...)
    this.counter++
    return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
    if this.counter == this.size {
        return false
    }
    this.queue = append(this.queue, value)
    this.counter++
    return true
}

func (this *MyCircularDeque) DeleteFront() bool {
    if this.counter == 0 {
        return false
    }
    this.queue = this.queue[1:]
    this.counter--
    return true
}

func (this *MyCircularDeque) DeleteLast() bool {
    if this.counter == 0 {
        return false
    }
    this.queue = this.queue[0 : this.counter-1]
    this.counter--
    return true
}

func (this *MyCircularDeque) GetFront() int {
    if this.counter == 0 {
        return -1
    }
    return this.queue[0]
}

func (this *MyCircularDeque) GetRear() int {
    if this.counter == 0 {
        return -1
    }
    return this.queue[this.counter-1]
}

func (this *MyCircularDeque) IsEmpty() bool {
    return this.counter == 0
}

func (this *MyCircularDeque) IsFull() bool {
    return this.counter == this.size
}

func (this *MyCircularDeque) Display() {
    fmt.Println(this.queue)
}
