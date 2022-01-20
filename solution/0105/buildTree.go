package _0105

import "github.com/raw34/leetcode/runtime"

func buildTree(preorder []int, inorder []int) *runtime.TreeNode {
    if len(preorder) == 0 || len(inorder) == 0 {
        return nil
    }

    indexMap := make(map[int]int, 0)
    for i, v := range inorder {
        indexMap[v] = i
    }

    val := preorder[0]
    root := &runtime.TreeNode{Val: val}
    mid := indexMap[val]
    root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
    root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])

    return root
}
