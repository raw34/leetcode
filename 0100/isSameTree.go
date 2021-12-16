package _0100

import "github.com/raw34/leetcode/runtime"

func isSameTree(p *runtime.TreeNode, q *runtime.TreeNode) bool {
    if p == nil && q == nil {
        return true
    }

    if p == nil && q != nil || p != nil && q == nil {
        return false
    }

    queue1 := []*runtime.TreeNode{p}
    queue2 := []*runtime.TreeNode{q}
    for len(queue1) > 0 && len(queue2) > 0 {
        root1 := queue1[0]
        queue1 = queue1[1:]
        root2 := queue2[0]
        queue2 = queue2[1:]

        if root1.Val != root2.Val {
            return false
        }

        if root1.Left != nil && root2.Left == nil || root1.Left == nil && root2.Left != nil {
            return false
        }

        if root1.Right != nil && root2.Right == nil || root1.Right == nil && root2.Right != nil {
            return false
        }

        if root1.Left != nil {
            queue1 = append(queue1, root1.Left)
        }
        if root1.Right != nil {
            queue1 = append(queue1, root1.Right)
        }
        if root2.Left != nil {
            queue2 = append(queue2, root2.Left)
        }
        if root2.Right != nil {
            queue2 = append(queue2, root2.Right)
        }
    }

    if len(queue1) > 0 || len(queue2) > 0 {
        return false
    }

    return true
}
