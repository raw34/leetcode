package _0652

import (
    "fmt"
    "github.com/raw34/leetcode/runtime"
)

func findDuplicateSubtrees(root *runtime.TreeNode) []*runtime.TreeNode {
    res := make([]*runtime.TreeNode, 0)
    counters := map[string]int{}

    var dfs func(root *runtime.TreeNode) string
    dfs = func(root *runtime.TreeNode) string {
        if root == nil {
            return "null"
        }
        str := fmt.Sprintf("%d,%s,%s", root.Val, dfs(root.Left), dfs(root.Right))
        counters[str]++
        if counters[str] == 2 {
            res = append(res, root)
        }

        return str
    }
    dfs(root)

    return res
}
