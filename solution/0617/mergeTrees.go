package _0617

import "github.com/raw34/leetcode/runtime"

func mergeTrees(t1 *runtime.TreeNode, t2 *runtime.TreeNode) *runtime.TreeNode {
    if t1 == nil {
        return t2
    }
    if t2 == nil {
        return t1
    }
    t1.Val += t2.Val
    t1.Left = mergeTrees(t1.Left, t2.Left)
    t1.Right = mergeTrees(t1.Right, t2.Right)

    return t1
}
