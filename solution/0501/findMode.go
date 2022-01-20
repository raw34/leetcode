package _0501

import "github.com/raw34/leetcode/runtime"

func findMode(root *runtime.TreeNode) []int {
    res := make([]int, 0)
    prevVal := 0
    currCount := 0
    maxCount := 1
    var dfs func(root *runtime.TreeNode)
    dfs = func(root *runtime.TreeNode) {
        if root == nil {
            return
        }

        dfs(root.Left)
        if root.Val == prevVal {
            currCount++
        } else {
            currCount = 1
            prevVal = root.Val
        }
        if currCount > maxCount {
            res = make([]int, 0)
            maxCount = currCount
        }
        if currCount >= maxCount {
            res = append(res, root.Val)
        }
        dfs(root.Right)
    }
    dfs(root)

    return res
}
