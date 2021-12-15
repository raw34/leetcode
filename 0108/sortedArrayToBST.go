package _0108

import "github.com/raw34/leetcode/runtime"

func sortedArrayToBST(nums []int) *runtime.TreeNode {
    var dfs func(start, end int) *runtime.TreeNode
    dfs = func(start, end int) *runtime.TreeNode {
        if start > end {
            return nil
        }
        mid := (start + end + 1) / 2
        root := &runtime.TreeNode{Val: nums[mid]}
        root.Left = dfs(start, mid-1)
        root.Right = dfs(mid+1, end)
        return root
    }

    return dfs(0, len(nums)-1)
}
