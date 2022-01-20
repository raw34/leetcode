package _0450

import (
    "github.com/raw34/leetcode/runtime"
)

func deleteNode(root *runtime.TreeNode, key int) *runtime.TreeNode {
    if root == nil {
        return root
    }

    if root.Val < key {
        root.Right = deleteNode(root.Right, key)
    } else if root.Val > key {
        root.Left = deleteNode(root.Left, key)
    } else {
        if root.Left == nil {
            return root.Right
        }
        if root.Right == nil {
            return root.Left
        }
        minNode := getMin(root.Right)
        root.Val = minNode.Val
        root.Right = deleteNode(root.Right, minNode.Val)
    }

    return root
}

func getMin(node *runtime.TreeNode) *runtime.TreeNode {
    for node.Left != nil {
        node = node.Left
    }
    return node
}
