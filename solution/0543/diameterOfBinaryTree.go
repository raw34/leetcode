package _0543

import "github.com/raw34/leetcode/runtime"

func diameterOfBinaryTree(root *runtime.TreeNode) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }
    res := 0
    var dfs func(root *runtime.TreeNode) int
    dfs = func(root *runtime.TreeNode) int {
        if root == nil {
            return 0
        }

        left := dfs(root.Left)
        right := dfs(root.Right)
        res = max(res, left+right)

        return max(left, right) + 1
    }
    dfs(root)

    return res
}
