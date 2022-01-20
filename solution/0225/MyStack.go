package _0225

import (
    "container/list"
    "fmt"
)

type MyStack struct {
    size   int
    queue1 list.List
    queue2 list.List
}

func Constructor() MyStack {
    return MyStack{}
}

func (this *MyStack) Push(x int) {
    this.queue1.PushBack(x)
    this.queue2.PushBack(x)
    if this.size >= 1 {
        this.queue2.Remove(this.queue2.Front())
    }
    this.size++
}

func (this *MyStack) Pop() int {
    val := this.queue2.Front().Value.(int)
    this.queue1.Remove(this.queue1.Back())
    this.queue2.Remove(this.queue2.Front())
    if this.size > 1 {
        this.queue2.PushBack(this.queue1.Back().Value)
    }
    this.size--
    return val
}

func (this *MyStack) Top() int {
    return this.queue2.Front().Value.(int)
}

func (this *MyStack) Empty() bool {
    return this.size == 0
}

func (this *MyStack) Display() {
    fmt.Print("stack1 = ")
    curr1 := this.queue1.Front()
    for curr1 != nil {
        fmt.Print(curr1.Value, " ")
        curr1 = curr1.Next()
    }
    fmt.Println()
    fmt.Print("stack2 = ")
    curr2 := this.queue2.Front()
    for curr2 != nil {
        fmt.Print(curr2.Value, " ")
        curr2 = curr2.Next()
    }
    fmt.Println()
}
