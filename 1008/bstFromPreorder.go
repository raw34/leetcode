package _1008

import "github.com/raw34/leetcode/runtime"

func bstFromPreorder(preorder []int) *runtime.TreeNode {
    var dfs func(start, end int) *runtime.TreeNode
    dfs = func(start, end int) *runtime.TreeNode {
        if start > end {
            return nil
        }

        val := preorder[0]
        root := &runtime.TreeNode{Val: val}

        // 寻找左右子树中间节点
        mid := end + 1
        for i := 0; i < len(preorder); i++ {
            if preorder[i] > val {
                mid = i
                break
            }
        }
        root.Left = bstFromPreorder(preorder[start+1 : mid])
        root.Right = bstFromPreorder(preorder[mid:])

        return root
    }

    return dfs(0, len(preorder)-1)
}
