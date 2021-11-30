package _110

import (
	"github.com/raw34/leetcode/runtime"
)

func isBalanced(root *runtime.TreeNode) bool {
	if root == nil {
		return true
	}

	return runtime.Abs(height(root.Left)-height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *runtime.TreeNode) int {
	if root == nil {
		return 0
	}

	return int(runtime.Max(height(root.Left), height(root.Right)) + 1)
}