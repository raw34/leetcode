package _0104

import (
    "github.com/raw34/leetcode/runtime"
)

func maxDepth(root *runtime.TreeNode) int {
    if root == nil {
        return 0
    }

    return int(runtime.Max(maxDepth(root.Left), maxDepth(root.Right))) + 1
}
