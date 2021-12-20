package _0113

import "github.com/raw34/leetcode/runtime"

func pathSum(root *runtime.TreeNode, targetSum int) [][]int {
    res := make([][]int, 0)
    path := make([]int, 0)
    var dfs func(root *runtime.TreeNode, path []int, sum int)
    dfs = func(root *runtime.TreeNode, path []int, sum int) {
        if root == nil {
            return
        }
        sum += root.Val
        path = append(path, root.Val)
        if root.Left == nil && root.Right == nil && sum == targetSum {
            res = append(res, append([]int{}, path...))
            return
        }
        dfs(root.Left, path, sum)
        dfs(root.Right, path, sum)
        path = path[:len(path)-1]
    }
    dfs(root, path, 0)

    return res
}
