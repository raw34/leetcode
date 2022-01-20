package _0460

import (
    "container/list"
)

type Node struct {
    key   int
    value int
    count int
}

type LFUCache struct {
    capacity int
    minCount int
    keyMap   map[int]*list.Element
    countMap map[int]*list.List
}

func Constructor(capacity int) LFUCache {
    return LFUCache{
        capacity: capacity,
        keyMap:   make(map[int]*list.Element, capacity),
        countMap: make(map[int]*list.List, capacity),
    }
}

func (this *LFUCache) Get(key int) int {
    if this.capacity == 0 {
        return -1
    }
    e, ok := this.keyMap[key]
    if !ok {
        return -1
    }
    this.update(e)
    node := e.Value.(*Node)
    return node.value
}

func (this *LFUCache) update(e *list.Element) {
    node := e.Value.(*Node)
    curList, ok := this.countMap[node.count]
    if !ok {
        return
    }
    curList.Remove(e)
    if curList.Len() == 0 {
        if node.count == this.minCount {
            this.minCount++
        }
    }
    node.count++
    newList, ok := this.countMap[node.count]
    if !ok {
        newList = list.New()
        this.countMap[node.count] = newList
    }
    newEle := newList.PushBack(node)
    this.keyMap[node.key] = newEle
}

func (this *LFUCache) Put(key int, value int) {
    minEle, ok := this.keyMap[key]
    if ok {
        node := minEle.Value.(*Node)
        node.value = value
        this.update(minEle)
        return
    }
    if len(this.keyMap) == this.capacity {
        minList, ok := this.countMap[this.minCount]
        if !ok {
            return
        }
        minEle = minList.Front()
        minList.Remove(minEle)
        minNode := minEle.Value.(*Node)
        delete(this.keyMap, minNode.key)
    }
    newNode := &Node{key, value, 1}
    newList, ok := this.countMap[newNode.count]
    if !ok {
        newList = list.New()
        this.countMap[newNode.count] = newList
    }
    newEle := newList.PushBack(newNode)
    this.keyMap[key] = newEle
    this.minCount = newNode.count
}
