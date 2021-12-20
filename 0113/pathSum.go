package _0113

import "github.com/raw34/leetcode/runtime"

func pathSum(root *runtime.TreeNode, targetSum int) [][]int {
    res := make([][]int, 0)
    path := make([]int, 0)
    var dfs func(root *runtime.TreeNode, num int)
    dfs = func(root *runtime.TreeNode, num int) {
        if root == nil {
            return
        }
        num -= root.Val
        path = append(path, root.Val)
        defer func() { path = path[:len(path)-1] }()
        if root.Left == nil && root.Right == nil && num == 0 {
            res = append(res, append([]int{}, path...))
            return
        }
        dfs(root.Left, num)
        dfs(root.Right, num)
    }
    dfs(root, targetSum)

    return res
}
