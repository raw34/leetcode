package _0101

import (
    "github.com/raw34/leetcode/runtime"
    "math"
)

func isSymmetric(root *runtime.TreeNode) bool {
    var dfs func(node1 *runtime.TreeNode, node2 *runtime.TreeNode) bool
    dfs = func(node1 *runtime.TreeNode, node2 *runtime.TreeNode) bool {
        if node1 == nil && node2 == nil {
            return true
        }
        if node1 == nil || node2 == nil {
            return false
        }
        return node1.Val == node2.Val && dfs(node1.Left, node2.Right) && dfs(node1.Right, node2.Left)
    }
    return dfs(root.Left, root.Right)
}

func isSymmetric2(root *runtime.TreeNode) bool {
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
