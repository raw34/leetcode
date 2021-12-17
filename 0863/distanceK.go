package _0863

import "github.com/raw34/leetcode/runtime"

func distanceK(root *runtime.TreeNode, target *runtime.TreeNode, k int) []int {
    parents := map[int]*runtime.TreeNode{}
    var findParent func(root *runtime.TreeNode)
    findParent = func(root *runtime.TreeNode) {
        if root == nil {
            return
        }
        if root.Left != nil {
            parents[root.Left.Val] = root
        }
        if root.Right != nil {
            parents[root.Right.Val] = root
        }

        findParent(root.Left)
        findParent(root.Right)
    }
    findParent(root)

    res := make([]int, 0)
    var dfs func(root, from *runtime.TreeNode, k int)
    dfs = func(root, from *runtime.TreeNode, k int) {
        if k == 0 {
            res = append(res, root.Val)
        }
        if root.Left != nil && root.Left != from {
            dfs(root.Left, root, k-1)
        }
        if root.Right != nil && root.Right != from {
            dfs(root.Right, root, k-1)
        }
        if parents[root.Val] != nil && parents[root.Val] != from {
            dfs(parents[root.Val], root, k-1)
        }
    }
    dfs(target, nil, k)

    return res
}
