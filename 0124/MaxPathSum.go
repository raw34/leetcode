package _124

import (
	"github.com/raw34/leetcode/runtime"
	"math"
)

func maxPathSum(root *runtime.TreeNode) int {
	maxSum := math.MinInt32
	var maxGain func(node *runtime.TreeNode) int
	maxGain = func(node *runtime.TreeNode) int {
		if node == nil {
			return 0
		}

		leftGain := runtime.Max(maxGain(node.Left), 0)
		rightGain := runtime.Max(maxGain(node.Right), 0)

		currSum := node.Val + leftGain + rightGain
		maxSum = runtime.Max(maxSum, currSum)

		return node.Val + runtime.Max(leftGain, rightGain)
	}
	maxGain(root)
	return maxSum
}