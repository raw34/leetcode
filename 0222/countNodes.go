package _0222

import "github.com/raw34/leetcode/runtime"

func countNodes(root *runtime.TreeNode) int {
    res := 0
    var dfs func(root *runtime.TreeNode)
    dfs = func(root *runtime.TreeNode) {
        if root == nil {
            return
        }
        dfs(root.Left)
        res++
        dfs(root.Right)
    }
    dfs(root)

    return res
}
