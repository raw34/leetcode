package _0101

import (
    "github.com/raw34/leetcode/runtime"
    "math"
)

func isSymmetric(root *runtime.TreeNode) bool {
    preorder := make([]int, 0)
    postorder := make([]int, 0)
    var dfs func(root *runtime.TreeNode)
    dfs = func(root *runtime.TreeNode) {
        if root == nil {
            preorder = append(preorder, math.MinInt32)
            postorder = append(postorder, math.MinInt32)
            return
        }
        preorder = append(preorder, root.Val)
        dfs(root.Left)
        dfs(root.Right)
        postorder = append(postorder, root.Val)
    }
    dfs(root)

    n := len(preorder)
    for i, j := 0, n-1; i < n; i, j = i+1, j-1 {
        if preorder[i] != postorder[j] {
            return false
        }
    }

    return true
}
