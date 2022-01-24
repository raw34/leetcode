package _0226

import "github.com/raw34/leetcode/runtime"

func invertTree(root *runtime.TreeNode) *runtime.TreeNode {
    var dfs func(root *runtime.TreeNode)
    dfs = func(root *runtime.TreeNode) {
        if root == nil {
            return
        }
        dfs(root.Left)
        dfs(root.Right)
        root.Left, root.Right = root.Right, root.Left
    }
    dfs(root)

    return root
}
