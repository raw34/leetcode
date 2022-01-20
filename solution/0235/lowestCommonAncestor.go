package _0235

import "github.com/raw34/leetcode/runtime"

func lowestCommonAncestor(root, p, q *runtime.TreeNode) *runtime.TreeNode {
    res := root
    for {
        if p.Val < res.Val && q.Val < res.Val {
            res = res.Left
        } else if p.Val > res.Val && q.Val > res.Val {
            res = res.Right
        } else {
            break
        }
    }

    return res
}
