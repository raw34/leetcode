package _0155

type MinStack struct {
    stack    []int
    minStack []int
}

func Constructor() MinStack {
    return MinStack{
        stack:    make([]int, 0),
        minStack: make([]int, 0),
    }
}

func (this *MinStack) Push(val int) {
    this.stack = append(this.stack, val)

    if len(this.minStack) == 0 || val <= this.GetMin() {
        this.minStack = append(this.minStack, val)
    }
}

func (this *MinStack) Pop() {
    val := this.stack[len(this.stack)-1]
    this.stack = this.stack[:len(this.stack)-1]

    if val == this.minStack[len(this.minStack)-1] {
        this.minStack = this.minStack[:len(this.minStack)-1]
    }
}

func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
    return this.minStack[len(this.minStack)-1]
}
