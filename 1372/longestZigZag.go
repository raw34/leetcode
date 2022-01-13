package _1372

import "github.com/raw34/leetcode/runtime"

func longestZigZag(root *runtime.TreeNode) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }
    res := 0
    var dfs func(root *runtime.TreeNode, parent string, length int)
    dfs = func(root *runtime.TreeNode, parent string, length int) {
        if root == nil {
            // 遍历到叶子节点时重新计算最长路径
            res = max(res, length-1)
            return
        }
        if parent == "left" {
            // 上一次访问左节点时，本次先访问右节点
            dfs(root.Right, "right", length+1)
            dfs(root.Left, "left", 1)
        } else {
            // 上一次访问右节点时，本次先访问左节点
            dfs(root.Left, "left", length+1)
            dfs(root.Right, "right", 1)
        }
    }
    dfs(root, "left", 0)
    dfs(root, "right", 0)

    return res
}
