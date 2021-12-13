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

func buildTree2(preorder []int, inorder []int) *runtime.TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    root := &runtime.TreeNode{Val: preorder[0]}
    stack := make([]*runtime.TreeNode, 0)
    stack = append(stack, root)
    var inorderIndex int
    for i := 1; i < len(preorder); i++ {
        preorderVal := preorder[i]
        node := stack[len(stack)-1]
        if node.Val != inorder[inorderIndex] {
            node.Left = &runtime.TreeNode{Val: preorderVal}
            stack = append(stack, node.Left)
        } else {
            for len(stack) != 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
                node = stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                inorderIndex++
            }
            node.Right = &runtime.TreeNode{Val: preorderVal}
            stack = append(stack, node.Right)
        }
    }
    return root
}
