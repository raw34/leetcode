package _1206

import (
    "math/rand"
    "time"
)

const MaxLevel = 32
const Probability = 0.25

type SkiplistNode struct {
    Val  int
    Next []*SkiplistNode
}

func NewSkiplistNode(val, level int) *SkiplistNode {
    return &SkiplistNode{Val: val, Next: make([]*SkiplistNode, level)}
}

type Skiplist struct {
    level int
    head  *SkiplistNode
}

func Constructor() Skiplist {
    return Skiplist{level: 1, head: NewSkiplistNode(0, MaxLevel)}
}

func (this *Skiplist) Search(target int) bool {
    prevNode := this.head
    // 逐层遍历，寻找前序节点
    for i := this.level - 1; i >= 0; i-- {
        prevNode = this.findPrevNode(prevNode, i, target)
        if prevNode.Next[i] != nil && prevNode.Next[i].Val == target {
            return true
        }
    }

    return false
}

func (this *Skiplist) findPrevNode(node *SkiplistNode, level, val int) *SkiplistNode {
    // 当前节点有后序节点，且节点值小于目标值，从下个节点找起
    for node.Next[level] != nil && node.Next[level].Val < val {
        node = node.Next[level]
    }

    return node
}
func (this *Skiplist) Add(num int) {
    newLevel := this.randomLevel()
    newNode := NewSkiplistNode(num, newLevel)
    prevNode := this.head
    for i := newLevel - 1; i >= 0; i-- {
        // 逐层遍历，寻找前序节点
        prevNode = this.findPrevNode(prevNode, i, num)
        if prevNode.Next[i] == nil {
            // 最近节点无后序节点，直接更新到该节点后
            prevNode.Next[i] = newNode
        } else {
            // 最近节点有后序节点，插入当前节点后，并将后序节点后移
            prevNode.Next[i], newNode.Next[i] = newNode, prevNode.Next[i]
        }
    }
    if newLevel > this.level {
        // 如果随机层比当前最大层大
        // 将新节点更新到每层第一个节点后
        for i := this.level; i < newLevel; i++ {
            this.head.Next[i] = newNode
        }
        this.level = newLevel
    }
}

func (this *Skiplist) randomLevel() int {
    level := 1
    rand.Seed(time.Now().UnixNano())
    for rand.Float32() < Probability && level < MaxLevel {
        level++
    }

    return level
}

func (this *Skiplist) Erase(num int) bool {
    res := false
    prevNode := this.head
    // 逐层遍历，寻找前序节点
    for i := this.level - 1; i >= 0; i-- {
        prevNode = this.findPrevNode(prevNode, i, num)
        // 当找到前序节点，且其后序节点为目标值，删除节点，后序节点前移
        if prevNode.Next[i] != nil && prevNode.Next[i].Val == num {
            prevNode.Next[i] = prevNode.Next[i].Next[i]
            res = true
            continue
        }
    }

    return res
}
