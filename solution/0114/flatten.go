package _0114

import "github.com/raw34/leetcode/runtime"

func flatten(root *runtime.TreeNode) {
    if root == nil {
        return
    }
    if root.Left != nil {
        left := root.Left
        curr := left
        for curr.Right != nil {
            curr = curr.Right
        }
        curr.Right = root.Right
        root.Right = left
        root.Left = nil
    }
    flatten(root.Right)
}
