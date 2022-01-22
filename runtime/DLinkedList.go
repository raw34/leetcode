package runtime

type DListNode struct {
    Val  int
    Prev *DListNode
    Next *DListNode
}

type DLinkedList struct {
    Head *DListNode
    Tail *DListNode
}
