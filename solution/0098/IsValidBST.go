package _0098

import (
    "github.com/raw34/leetcode/runtime"
    "math"
)

func isValidBST(root *runtime.TreeNode) bool {
    if root == nil {
        return true
    }

    return validate(root, math.MinInt64, math.MaxInt64)
}

func validate(root *runtime.TreeNode, lower, upper int) bool {
    if root == nil {
        return true
    }

    if root.Val <= lower || root.Val >= upper {
        return false
    }

    return validate(root.Left, lower, root.Val) && validate(root.Right, root.Val, upper)
}
