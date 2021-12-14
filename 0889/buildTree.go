package _0889

import "github.com/raw34/leetcode/runtime"

func constructFromPrePost(preorder []int, postorder []int) *runtime.TreeNode {
    if len(preorder) == 0 || len(postorder) == 0 {
        return nil
    }

    val := preorder[0]
    root := &runtime.TreeNode{Val: val}
    if len(postorder) == 1 {
        return root
    }

    mid := 0
    for i := 0; i < len(postorder); i++ {
        if postorder[i] == preorder[1] {
            mid = i + 1
            break
        }
    }

    root.Left = constructFromPrePost(preorder[1:mid+1], postorder[0:mid])
    root.Right = constructFromPrePost(preorder[mid+1:], postorder[mid:len(postorder)-1])

    return root
}
