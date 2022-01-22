package runtime

import "fmt"

type DListNode struct {
    Val  int
    Prev *DListNode
    Next *DListNode
}

type DLinkedList struct {
    head *DListNode
    tail *DListNode
    size int
}

func NewDLinkedList() DLinkedList {
    head := &DListNode{}
    tail := &DListNode{}
    head.Next = tail
    tail.Prev = head
    return DLinkedList{head, tail, 0}
}

func (this *DLinkedList) Get(index int) int {
    if index < 0 || index >= this.size {
        return -1
    }

    curr := this.head
    for i := 0; i < index+1; i++ {
        curr = curr.Next
    }

    return curr.Val
}

func (this *DLinkedList) AddAtHead(val int) {
    prev := this.head
    next := prev.Next
    node := &DListNode{Val: val}
    node.Prev = prev
    node.Next = next
    prev.Next = node
    next.Prev = node
    this.size++
}

func (this *DLinkedList) AddAtTail(val int) {
    prev := this.tail.Prev
    next := prev.Next
    node := &DListNode{Val: val}
    node.Prev = prev
    node.Next = next
    prev.Next = node
    next.Prev = node
    this.size++
}

func (this *DLinkedList) AddAtIndex(index int, val int) {
    if index < 0 || index > this.size {
        return
    }

    prev := this.head
    for i := 0; i < index; i++ {
        prev = prev.Next
    }

    next := prev.Next
    node := &DListNode{Val: val}
    node.Prev = prev
    node.Next = next
    prev.Next = node
    next.Prev = node
    this.size++
}

func (this *DLinkedList) DeleteAtIndex(index int) {
    if index < 0 || index >= this.size {
        return
    }

    prev := this.head
    for i := 0; i < index; i++ {
        prev = prev.Next
    }

    next := prev.Next.Next
    prev.Next = next
    next.Prev = prev
    this.size--
}

func (this *DLinkedList) Display() {
    curr := this.head
    for curr != nil {
        fmt.Print(curr, " ")
        curr = curr.Next
    }
    fmt.Println()
}
