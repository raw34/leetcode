package _1206

import "math/rand"

const MaxLevel = 32
const P = 0.25

type SkiplistNode struct {
    Val  int
    Next []*SkiplistNode
}

type Skiplist struct {
    level int
    head  *SkiplistNode
}

func Constructor() Skiplist {
    head := &SkiplistNode{Next: make([]*SkiplistNode, MaxLevel)}
    return Skiplist{level: 1, head: head}
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
    level := this.randomLevel()
    newNode := &SkiplistNode{Val: num, Next: make([]*SkiplistNode, level)}
    if level > this.level {
        // 如果随机层比当前最大层大
        // 将新节点更新到每层第一个节点后
        for i := this.level; i < level; i++ {
            this.head.Next[i] = newNode
        }
        this.level = level
    } else {
        // 如果随机层比当前最大层小
        // 逐层遍历，寻找前序节点
        prevNode := this.head
        for i := level - 1; i >= 0; i-- {
            prevNode = this.findPrevNode(prevNode, i, num)
            if prevNode.Next[i] == nil {
                // 最近节点无后序节点，直接更新到该节点后
                prevNode.Next[i] = newNode
            } else {
                // 最近节点有后序节点，插入当前节点后，并将后序节点后移
                prevNode.Next[i], newNode.Next[i] = newNode, prevNode.Next[i]
            }
        }
    }
}

func (this *Skiplist) randomLevel() int {
    level := 1
    for rand.Float64() < P && level < MaxLevel {
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
