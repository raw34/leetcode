package _0404

import "github.com/raw34/leetcode/runtime"

func sumOfLeftLeaves(root *runtime.TreeNode) int {
    isLeafNode := func(node *runtime.TreeNode) bool {
        return node.Left == nil && node.Right == nil
    }

    res := 0
    queue := []*runtime.TreeNode{root}
    for len(queue) > 0 {
        root = queue[0]
        queue = queue[1:]
        if root.Left != nil {
            if isLeafNode(root.Left) {
                res += root.Left.Val
            }
            queue = append(queue, root.Left)
        }
        if root.Right != nil && !isLeafNode(root.Right) {
            queue = append(queue, root.Right)
        }
    }

    return res
}
