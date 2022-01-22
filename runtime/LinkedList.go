package runtime

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

type LinkedList struct {
    head *ListNode
    size int
}

func NewLinkedList() LinkedList {
    return LinkedList{&ListNode{}, 0}
}

func (this *LinkedList) Get(index int) int {
    if index < 0 || index >= this.size {
        return -1
    }

    curr := this.head
    for i := 0; i < index+1; i++ {
        curr = curr.Next
    }

    return curr.Val
}

func (this *LinkedList) AddAtHead(val int) {
    this.AddAtIndex(0, val)
}

func (this *LinkedList) AddAtTail(val int) {
    this.AddAtIndex(this.size, val)
}

func (this *LinkedList) AddAtIndex(index int, val int) {
    if index < 0 || index > this.size {
        return
    }

    curr := this.head
    for i := 0; i < index; i++ {
        curr = curr.Next
    }

    node := &ListNode{Val: val}
    node.Next = curr.Next
    curr.Next = node
    this.size++
}

func (this *LinkedList) DeleteAtIndex(index int) {
    if index < 0 || index >= this.size {
        return
    }

    curr := this.head
    for i := 0; i < index; i++ {
        curr = curr.Next
    }

    curr.Next = curr.Next.Next
    this.size--
}

func (this *LinkedList) Display() {
    curr := this.head
    i := 0
    for curr != nil {
        fmt.Print(curr.Val, " ")
        curr = curr.Next
        i++
    }
    fmt.Println()
}
