package _0129

import "github.com/raw34/leetcode/runtime"

func sumNumbers(root *runtime.TreeNode) int {
    var dfs func(root *runtime.TreeNode, sum int) int
    dfs = func(root *runtime.TreeNode, sum int) int {
        if root == nil {
            return 0
        }
        sum = sum*10 + root.Val
        if root.Left == nil && root.Right == nil {
            return sum
        }

        return dfs(root.Left, sum) + dfs(root.Right, sum)
    }

    return dfs(root, 0)
}
