package _110

import (
	"github.com/raw34/leetcode/runtime"
	"math"
)

func isBalanced(root *runtime.TreeNode) bool {
	if root == nil {
		return true
	}

	return math.Abs(float64(height(root.Left)-height(root.Right))) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *runtime.TreeNode) int {
	if root == nil {
		return 0
	}

	return int(math.Max(float64(height(root.Left)), float64(height(root.Right))) + 1)
}