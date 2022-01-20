package _0099

import (
    "github.com/raw34/leetcode/runtime"
    "math"
)

func recoverTree(root *runtime.TreeNode) {
    var left, right *runtime.TreeNode
    prev := &runtime.TreeNode{Val: math.MinInt32}
    var dfs func(root *runtime.TreeNode)
    dfs = func(root *runtime.TreeNode) {
        if root == nil {
            return
        }
        dfs(root.Left)
        if left == nil && prev.Val > root.Val {
            left = prev
        }
        if left != nil && prev.Val > root.Val {
            right = root
        }
        prev = root
        dfs(root.Right)
    }
    dfs(root)
    left.Val, right.Val = right.Val, left.Val
}
