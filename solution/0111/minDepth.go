package _0111

import (
    "github.com/raw34/leetcode/runtime"
    "math"
)

func minDepth(root *runtime.TreeNode) int {
    min := func(a, b int) int {
        if a > b {
            return b
        }
        return a
    }
    res := math.MaxInt32
    var dfs func(root *runtime.TreeNode, depth int)
    dfs = func(root *runtime.TreeNode, depth int) {
        if root == nil {
            return
        }
        if root.Left == nil && root.Right == nil {
            res = min(res, depth+1)
            return
        }
        if root.Left != nil {
            dfs(root.Left, depth+1)
        }
        if root.Right != nil {
            dfs(root.Right, depth+1)
        }
    }
    dfs(root, 0)
    if res == math.MaxInt32 {
        res = 0
    }

    return res
}
