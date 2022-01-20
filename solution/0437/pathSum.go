package _0437

import "github.com/raw34/leetcode/runtime"

func pathSum(root *runtime.TreeNode, targetSum int) int {
    res := 0
    if root == nil {
        return res
    }

    var dfs2 func(root *runtime.TreeNode, sum int)
    dfs2 = func(root *runtime.TreeNode, sum int) {
        if sum == targetSum {
            res++
        }

        if root.Left != nil {
            dfs2(root.Left, sum+root.Left.Val)
        }

        if root.Right != nil {
            dfs2(root.Right, sum+root.Right.Val)
        }
    }

    var dfs1 func(root *runtime.TreeNode)
    dfs1 = func(root *runtime.TreeNode) {
        if root == nil {
            return
        }
        dfs2(root, root.Val)
        dfs1(root.Left)
        dfs1(root.Right)
    }
    dfs1(root)

    return res
}
