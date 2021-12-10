package _0707

import (
    "fmt"
)

type ListNode struct {
    Val  int
    Next *ListNode
}

type MyLinkedList struct {
    size int
    Head *ListNode
}

func Constructor() MyLinkedList {
    head := &ListNode{Val: 0}
    return MyLinkedList{0, head}
}

func (this *MyLinkedList) Get(index int) int {
    if index < 0 || index >= this.size {
        return -1
    }

    curr := this.Head
    for i := 0; i <= index; i++ {
        curr = curr.Next
    }

    return curr.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
    this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
    this.AddAtIndex(this.size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
    if index < 0 || index > this.size {
        return
    }

    curr := this.Head
    for i := 0; i < index; i++ {
        curr = curr.Next
    }

    node := &ListNode{Val: val}
    node.Next = curr.Next
    curr.Next = node
    this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
    if index < 0 || index >= this.size {
        return
    }

    curr := this.Head
    for i := 0; i < index; i++ {
        curr = curr.Next
    }

    curr.Next = curr.Next.Next
    this.size--
}

func (this *MyLinkedList) Display() {
    curr := this.Head
    i := 0
    for curr != nil {
        fmt.Print(curr.Val, " ")
        curr = curr.Next
        i++
    }
    fmt.Println()
}
