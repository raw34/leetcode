package _0538

import "github.com/raw34/leetcode/runtime"

func convertBST(root *runtime.TreeNode) *runtime.TreeNode {
    nodes := make([]*runtime.TreeNode, 0)
    var dfs func(root *runtime.TreeNode)
    dfs = func(root *runtime.TreeNode) {
        if root == nil {
            return
        }
        dfs(root.Left)
        nodes = append(nodes, root)
        dfs(root.Right)
    }
    dfs(root)

    for i := len(nodes) - 2; i >= 0; i-- {
        nodes[i].Val += nodes[i+1].Val
    }

    return root
}
