package _0105

import "github.com/raw34/leetcode/runtime"

func buildTree(preorder []int, inorder []int) *runtime.TreeNode {
    indexMap := make(map[int]int, 0)
    for i, v := range inorder {
        indexMap[v] = i
    }

    var dfs func(preorderLeft, preorderRight, inorderLeft, inorderRight int) *runtime.TreeNode
    dfs = func(preorderLeft, preorderRight, inorderLeft, inorderRight int) *runtime.TreeNode {
        if preorderLeft > preorderRight {
            return nil
        }
        preorderRoot := preorderLeft
        val := preorder[preorderRoot]
        inorderRoot := indexMap[val]
        leftTreeSize := inorderRoot - inorderLeft

        root := &runtime.TreeNode{Val: val}
        root.Left = dfs(preorderLeft+1, preorderLeft+leftTreeSize, inorderLeft, inorderRoot-1)
        root.Right = dfs(preorderLeft+leftTreeSize+1, preorderRight, inorderRoot+1, inorderRight)

        return root
    }

    n := len(preorder)
    return dfs(0, n-1, 0, n-1)
}
