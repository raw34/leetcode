package _124

import (
	"github.com/raw34/leetcode/runtime"
	"math"
)

func maxPathSum(root *runtime.TreeNode) int {
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	maxSum := math.MinInt32
	var maxGain func(node *runtime.TreeNode) int
	maxGain = func(node *runtime.TreeNode) int {
		if node == nil {
			return 0
		}

		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)

		currSum := node.Val + leftGain + rightGain
		maxSum = max(maxSum, currSum)

		return node.Val + max(leftGain, rightGain)
	}
	maxGain(root)
	return maxSum
}