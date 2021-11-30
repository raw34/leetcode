package _102

import "github.com/raw34/leetcode/runtime"

func levelOrder(root *runtime.TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := []*runtime.TreeNode{root}
	for level := 0; len(queue) > 0; level++ {
		length := len(queue)
		res = append(res, []int{})
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			res[level] = append(res[level], node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return res
}