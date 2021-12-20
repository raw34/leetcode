package _0113

import "github.com/raw34/leetcode/runtime"

func pathSum(root *runtime.TreeNode, targetSum int) [][]int {
    res := make([][]int, 0)
    path := make([]int, 0)
    var dfs func(root *runtime.TreeNode, path []int, num int)
    dfs = func(root *runtime.TreeNode, path []int, num int) {
        if root == nil {
            return
        }
        num += root.Val
        path = append(path, root.Val)
        if root.Left == nil && root.Right == nil && num == targetSum {
            res = append(res, append([]int{}, path...))
            return
        }
        dfs(root.Left, path, num)
        dfs(root.Right, path, num)
        path = path[:len(path)-1]
    }
    dfs(root, path, 0)

    return res
}
