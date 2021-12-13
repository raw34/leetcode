package _0144

import "github.com/raw34/leetcode/runtime"

func preorderTraversal(root *runtime.TreeNode) []int {
    stack := make([]int, 0)

    var dfs func(root *runtime.TreeNode)
    dfs = func(root *runtime.TreeNode) {
        if root == nil {
            return
        }
        stack = append(stack, root.Val)
        dfs(root.Left)
        dfs(root.Right)
    }
    dfs(root)

    return stack
}
