package _107

import "github.com/raw34/leetcode/runtime"

func levelOrderBottom(root *runtime.TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := []*runtime.TreeNode{root}
	for i := 0; len(queue) > 0; i++ {
		length := len(queue)
		curr := make([][]int, 1)
		for j := 0; j < length; j++ {
			node := queue[0]
			queue = queue[1:]
			curr[0] = append(curr[0], node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(curr, res...)
	}

	return res
}