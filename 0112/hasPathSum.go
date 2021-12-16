package _0112

import "github.com/raw34/leetcode/runtime"

func hasPathSum(root *runtime.TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }
    if root.Left == nil && root.Right == nil {
        return root.Val == targetSum
    }

    return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}
