package _0572

import "github.com/raw34/leetcode/runtime"

func isSubtree(root *runtime.TreeNode, subRoot *runtime.TreeNode) bool {
    if root == nil {
        return false
    }

    var check func(a, b *runtime.TreeNode) bool
    check = func(a, b *runtime.TreeNode) bool {
        if a == nil && b == nil {
            return true
        }
        if a == nil || b == nil || a.Val != b.Val {
            return false
        }
        return check(a.Left, b.Left) && check(a.Right, b.Right)
    }

    return check(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}
