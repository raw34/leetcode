package _0108

import "github.com/raw34/leetcode/runtime"

func sortedArrayToBST(nums []int) *runtime.TreeNode {
    var dfs func(nums []int, left, right int) *runtime.TreeNode
    dfs = func(nums []int, left, right int) *runtime.TreeNode {
        if left > right {
            return nil
        }
        mid := (left + right) / 2
        root := &runtime.TreeNode{Val: nums[mid]}
        root.Left = dfs(nums, left, mid-1)
        root.Right = dfs(nums, mid+1, right)
        return root
    }

    return dfs(nums, 0, len(nums)-1)
}
