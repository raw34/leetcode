package _0173

import (
    "github.com/raw34/leetcode/runtime"
)

type BSTIterator struct {
    queue []int
}

func Constructor(root *runtime.TreeNode) BSTIterator {
    it := BSTIterator{}
    it.inorder(root)

    return it
}

func (this *BSTIterator) inorder(root *runtime.TreeNode) {
    if root == nil {
        return
    }
    if root.Left != nil {
        this.inorder(root.Left)
    }
    this.queue = append(this.queue, root.Val)
    if root.Right != nil {
        this.inorder(root.Right)
    }
}

func (this *BSTIterator) Next() int {
    val := this.queue[0]
    this.queue = this.queue[1:]
    return val
}

func (this *BSTIterator) HasNext() bool {
    return len(this.queue) > 0
}
