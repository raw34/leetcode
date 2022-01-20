package _0095

import "github.com/raw34/leetcode/runtime"

func generateTrees(n int) []*runtime.TreeNode {
    var helper func(start, end int) []*runtime.TreeNode
    helper = func(start, end int) []*runtime.TreeNode {
        if start > end {
            return []*runtime.TreeNode{nil}
        }

        trees := make([]*runtime.TreeNode, 0)
        for i := start; i <= end; i++ {
            leftTrees := helper(start, i-1)
            rightTrees := helper(i+1, end)
            for _, left := range leftTrees {
                for _, right := range rightTrees {
                    tree := &runtime.TreeNode{}
                    tree.Val = i
                    tree.Left = left
                    tree.Right = right
                    trees = append(trees, tree)
                }
            }
        }

        return trees
    }

    return helper(1, n)
}
