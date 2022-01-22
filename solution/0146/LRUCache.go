package _0146

import "fmt"

type DListNode struct {
    Key  int
    Val  int
    Prev *DListNode
    Next *DListNode
}

type DLinkedList struct {
    head *DListNode
    tail *DListNode
    size int
}

func (this *DLinkedList) AddToHead(node *DListNode) {
    node.Prev = this.head
    node.Next = this.head.Next
    this.head.Next.Prev = node
    this.head.Next = node
    this.size++
}

func (this *DLinkedList) RemoveNode(node *DListNode) {
    node.Prev.Next = node.Next
    node.Next.Prev = node.Prev
    this.size--
}

func (this *DLinkedList) RemoveTail() *DListNode {
    node := this.tail.Prev
    this.RemoveNode(node)
    return node
}

func (this *DLinkedList) Display() {
    curr := this.head
    for curr != nil {
        fmt.Print(curr, " ")
        curr = curr.Next
    }
    fmt.Println()
}

type LRUCache struct {
    capacity int
    hash     map[int]*DListNode
    list     *DLinkedList
}

func Constructor(capacity int) LRUCache {
    head := &DListNode{}
    tail := &DListNode{}
    head.Next = tail
    tail.Prev = head

    return LRUCache{
        capacity: capacity,
        hash:     map[int]*DListNode{},
        list:     &DLinkedList{head, tail, 0},
    }
}

func (this *LRUCache) Get(key int) int {
    node, ok := this.hash[key]
    if !ok {
        // 未命中缓存，直接返回
        return -1
    }
    // 更新缓存热度
    this.list.RemoveNode(node)
    this.list.AddToHead(node)
    // 返回元素
    return node.Val
}

func (this *LRUCache) Put(key int, value int) {
    if node, ok := this.hash[key]; ok {
        // 如果key存在，删除元素
        this.list.RemoveNode(node)
    }
    if this.list.size == this.capacity {
        // 如果空间已满，清理缓存
        removed := this.list.RemoveTail()
        delete(this.hash, removed.Key)
    }
    // 添加缓存和度
    node := &DListNode{Key: key, Val: value}
    this.hash[key] = node
    this.list.AddToHead(node)
}
