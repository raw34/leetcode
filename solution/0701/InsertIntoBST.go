package _0701

import "github.com/raw34/leetcode/runtime"

func insertIntoBST(root *runtime.TreeNode, val int) *runtime.TreeNode {
    if root == nil {
        return &runtime.TreeNode{Val: val}
    }

    node := root
    for node != nil {
        if val < node.Val {
            if node.Left == nil {
                node.Left = &runtime.TreeNode{Val: val}
                break
            }
            node = node.Left
        } else {
            if node.Right == nil {
                node.Right = &runtime.TreeNode{Val: val}
                break
            }
            node = node.Right
        }
    }

    return root
}
