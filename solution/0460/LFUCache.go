package _0460

import (
    "fmt"
)

type DListNode struct {
    Key   int
    Val   int
    Count int
    Prev  *DListNode
    Next  *DListNode
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

type LFUCache struct {
    capacity int
    minCount int
    cache    map[int]*DListNode
    counter  map[int]*DLinkedList
}

func Constructor(capacity int) LFUCache {
    return LFUCache{
        capacity: capacity,
        cache:    map[int]*DListNode{},
        counter:  map[int]*DLinkedList{},
    }
}

func (this *LFUCache) Get(key int) int {
    node, ok := this.cache[key]
    if !ok || this.capacity == 0 {
        // 未命中，直接返回
        return -1
    }
    // 更新热度
    this.updateNodeCount(node)
    // 返回元素
    return node.Val
}

func (this *LFUCache) updateNodeCount(node *DListNode) {
    // 删除旧热度
    list := this.counter[node.Count]
    list.RemoveNode(node)
    if node.Count == this.minCount && list.size == 0 {
        // 更新热度最小值
        this.minCount++
    }
    // 更新新热度
    node.Count++
    this.getNodeList(node.Count).AddToHead(node)
}

func (this *LFUCache) getNodeList(count int) *DLinkedList {
    list, ok := this.counter[count]
    if !ok {
        list = &DLinkedList{&DListNode{}, &DListNode{}, 0}
        list.head.Next = list.tail
        list.tail.Prev = list.head
        this.counter[count] = list
    }

    return list
}

func (this *LFUCache) Put(key int, value int) {
    if this.capacity == 0 {
        return
    }
    // 如果key存在，更新缓存、更新热度
    if node, ok := this.cache[key]; ok {
        node.Val = value
        this.updateNodeCount(node)
        return
    }
    // 如果空间已满，清理缓存、清理热度
    if len(this.cache) == this.capacity {
        removed := this.counter[this.minCount].RemoveTail()
        delete(this.cache, removed.Key)
    }
    // 添加缓存、更新热度
    node := &DListNode{Key: key, Val: value, Count: 1}
    this.cache[key] = node
    this.getNodeList(node.Count).AddToHead(node)
    // 更新热度最小值
    this.minCount = node.Count
}
