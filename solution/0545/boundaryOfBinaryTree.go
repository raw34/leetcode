package _0545

import "github.com/raw34/leetcode/runtime"

func boundaryOfBinaryTree(root *runtime.TreeNode) []int {
    res := []int{root.Val}
    if root.Left == nil && root.Right == nil {
        return res
    }

    left := make([]int, 0)
    var dfsLeft func(root *runtime.TreeNode)
    dfsLeft = func(root *runtime.TreeNode) {
        if root != nil && (root.Left != nil || root.Right != nil) {
            left = append(left, root.Val)
            if root.Left != nil {
                dfsLeft(root.Left)
            } else {
                dfsLeft(root.Right)
            }
        }
    }
    dfsLeft(root.Left)

    bottom := make([]int, 0)
    var dfsBottom func(root *runtime.TreeNode)
    dfsBottom = func(root *runtime.TreeNode) {
        if root == nil {
            return
        }
        if root.Left == nil && root.Right == nil {
            bottom = append(bottom, root.Val)
        } else {
            dfsBottom(root.Left)
            dfsBottom(root.Right)
        }
    }
    dfsBottom(root)

    right := make([]int, 0)
    var dfsRight func(root *runtime.TreeNode)
    dfsRight = func(root *runtime.TreeNode) {
        if root != nil && (root.Left != nil || root.Right != nil) {
            right = append([]int{root.Val}, right...)
            if root.Right != nil {
                dfsRight(root.Right)
            } else {
                dfsRight(root.Left)
            }
        }
    }
    dfsRight(root.Right)

    res = append(res, append(left, append(bottom, right...)...)...)
    return res
}
