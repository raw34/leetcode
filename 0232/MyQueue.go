package _232

type MyQueue struct {
	stack1 []int
}

func Constructor() MyQueue {
	return MyQueue{
		stack1: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int)  {
	this.stack1 = append(this.stack1, x)
}

func (this *MyQueue) Pop() int {
	x := this.stack1[0]
	this.stack1 = this.stack1[1:]

	return x
}

func (this *MyQueue) Peek() int {
	return this.stack1[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.stack1) == 0
}
