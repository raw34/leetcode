package _0707

import (
    "fmt"
)

type ListNode struct {
    Val  int
    Next *ListNode
}

type MyLinkedList struct {
    Head *ListNode
}

func Constructor() MyLinkedList {
    return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
    curr := this.Head
    for i := 0; i < index; i++ {
        if curr == nil {
            break
        }
        curr = curr.Next
    }
    if curr == nil {
        return -1
    }

    return curr.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
    newHead := &ListNode{Val: val}
    newHead.Next = this.Head
    this.Head = newHead
}

func (this *MyLinkedList) AddAtTail(val int) {
    curr := this.Head
    for curr != nil && curr.Next != nil {
        curr = curr.Next
    }
    newTail := &ListNode{Val: val}
    if curr == nil {
        this.Head = newTail
    } else {
        curr.Next = newTail
    }
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
    curr := this.Head
    if curr == nil && index > 0 {
        return
    }

    for i := 0; i < index-1; i++ {
        curr = curr.Next
    }

    node := &ListNode{Val: val}
    if index == 0 || curr == nil {
        node.Next = curr
        this.Head = node
    } else {
        node.Next = curr.Next
        curr.Next = node
    }
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
    curr := this.Head
    for i := 0; i < index-1; i++ {
        curr = curr.Next
    }

    if curr == nil {
        return
    }

    if index == 0 {
        this.Head = curr.Next
    } else if curr.Next != nil {
        curr.Next = curr.Next.Next
    }
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
