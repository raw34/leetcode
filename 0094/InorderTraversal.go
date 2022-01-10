package _0094

import "github.com/raw34/leetcode/runtime"

var res []int

func inorderTraversal(root *runtime.TreeNode) []int {
    res = make([]int, 0)
    inorder(root)

    return res
}

func inorder(node *runtime.TreeNode) {
    if node == nil {
        return
    }

    inorder(node.Left)
    res = append(res, node.Val)
    inorder(node.Right)
}
