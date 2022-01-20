package _1008

import "github.com/raw34/leetcode/runtime"

func bstFromPreorder(preorder []int) *runtime.TreeNode {
    var dfs func(left, right int) *runtime.TreeNode
    dfs = func(left, right int) *runtime.TreeNode {
        if left > right {
            return nil
        }

        val := preorder[0]
        root := &runtime.TreeNode{Val: val}

        // 寻找左右子树中间节点
        mid := right + 1
        for i := 0; i < len(preorder); i++ {
            if preorder[i] > val {
                mid = i
                break
            }
        }
        root.Left = bstFromPreorder(preorder[left+1 : mid])
        root.Right = bstFromPreorder(preorder[mid:])

        return root
    }

    return dfs(0, len(preorder)-1)
}
