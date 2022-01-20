package _0257

import (
    "github.com/raw34/leetcode/runtime"
    "strconv"
)

func binaryTreePaths(root *runtime.TreeNode) []string {
    res := make([]string, 0)
    if root == nil {
        return res
    }
    var dfs func(root *runtime.TreeNode, path string)
    dfs = func(root *runtime.TreeNode, path string) {
        if root == nil {
            return
        }
        path += strconv.Itoa(root.Val)
        if root.Left == nil && root.Right == nil {
            res = append(res, path)
        } else {
            path += "->"
            dfs(root.Left, path)
            dfs(root.Right, path)
        }
    }
    dfs(root, "")

    return res
}
