package _0337

import "github.com/raw34/leetcode/runtime"

func rob(root *runtime.TreeNode) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }

    var dfs func(root *runtime.TreeNode) []int
    dfs = func(root *runtime.TreeNode) []int {
        // res[0]表示不抢当前节点时的最大收益
        // res[1]表示抢当前节点时的最大收益
        res := []int{0, 0}
        if root == nil {
            return res
        }
        left := dfs(root.Left)
        right := dfs(root.Right)
        // 不抢当前节点时，收益为前驱左右节点最大值之和
        res[0] = max(left[0], left[1]) + max(right[0], right[1])
        // 抢当前节点时，收益为当前节点与前驱左右节点都不抢情况之和
        res[1] = root.Val + left[0] + right[0]
        return res
    }
    res := dfs(root)

    return max(res[0], res[1])
}
