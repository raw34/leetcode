package _0285

import "github.com/raw34/leetcode/runtime"

func inorderSuccessor(root *runtime.TreeNode, p *runtime.TreeNode) *runtime.TreeNode {
    var res *runtime.TreeNode
    for root != nil {
        if root.Val <= p.Val {
            root = root.Right
        } else {
            res = root
            root = root.Left
        }
    }

    return res
}
