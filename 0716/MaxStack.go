package _716

import "fmt"

type MaxStack struct {
    stack1 []int
    stack2 []int
}

func Constructor() MaxStack {
    return MaxStack{}
}

func (this *MaxStack) Push(x int) {
    this.stack1 = append(this.stack1, x)
    if len(this.stack2) == 0 || this.PeekMax() <= x {
        this.stack2 = append(this.stack2, x)
    } else {
        this.stack2 = append(this.stack2, this.PeekMax())
    }
}

func (this *MaxStack) Pop() int {
    val := this.stack1[len(this.stack1)-1]
    this.stack1 = this.stack1[0 : len(this.stack1)-1]
    this.stack2 = this.stack2[0 : len(this.stack2)-1]
    return val
}

func (this *MaxStack) Top() int {
    return this.stack1[len(this.stack1)-1]
}

func (this *MaxStack) PeekMax() int {
    return this.stack2[len(this.stack2)-1]
}

func (this *MaxStack) PopMax() int {
    max := this.PeekMax()
    if max == this.stack1[len(this.stack1)-1] {
        this.stack1 = this.stack1[0 : len(this.stack1)-1]
        this.stack2 = this.stack2[0 : len(this.stack2)-1]
    } else {
        temp := make([]int, 0)
        for this.Top() != max {
            temp = append(temp, this.Pop())
        }
        this.Pop()
        for len(temp) > 0 {
            this.Push(temp[len(temp)-1])
            temp = temp[0 : len(temp)-1]
        }
    }
    return max
}

func (this *MaxStack) Display() {
    fmt.Println("q1", this.stack1, "q2", this.stack2)
}
