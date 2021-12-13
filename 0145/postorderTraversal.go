package _0145

import "github.com/raw34/leetcode/runtime"

func postorderTraversal(root *runtime.TreeNode) []int {
    stack := make([]int, 0)

    var dfs func(root *runtime.TreeNode)
    dfs = func(root *runtime.TreeNode) {
        if root == nil {
            return
        }
        dfs(root.Left)
        dfs(root.Right)
        stack = append(stack, root.Val)
    }
    dfs(root)

    return stack
}
