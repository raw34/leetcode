package _104

import (
	"github.com/raw34/leetcode/runtime"
	"math"
)

func maxDepth(root *runtime.TreeNode) int {
	if root == nil {
		return 0
	}

	return int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right)))) + 1
}