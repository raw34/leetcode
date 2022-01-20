package _0106

import "github.com/raw34/leetcode/runtime"

func buildTree(inorder []int, postorder []int) *runtime.TreeNode {
    if len(inorder) == 0 || len(postorder) == 0 {
        return nil
    }

    indexMap := make(map[int]int, 0)
    for i, v := range inorder {
        indexMap[v] = i
    }

    val := postorder[len(postorder)-1]
    mid := indexMap[val]
    root := &runtime.TreeNode{Val: val}
    root.Left = buildTree(inorder[:mid], postorder[:mid])
    root.Right = buildTree(inorder[mid+1:], postorder[mid:len(postorder)-1])

    return root
}
