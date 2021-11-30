package _236

import "github.com/raw34/leetcode/runtime"

func lowestCommonAncestor(root, p, q *runtime.TreeNode) *runtime.TreeNode {
	parents := make(map[int]*runtime.TreeNode, 0)
	visited := make(map[int]bool, 0)

	var dfs func(node *runtime.TreeNode)
	dfs = func(node *runtime.TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			parents[node.Left.Val] = node
			dfs(node.Left)
		}
		if node.Right != nil {
			parents[node.Right.Val] = node
			dfs(node.Right)
		}
	}
	dfs(root)

	for p != nil {
		visited[p.Val] = true
		p = parents[p.Val]
	}

	for q != nil {
		if visited[q.Val] {
			return q
		}
		q = parents[q.Val]
	}

	return nil
}