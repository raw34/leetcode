package _0863

import "github.com/raw34/leetcode/runtime"

func distanceK(root *runtime.TreeNode, target *runtime.TreeNode, k int) []int {
    // 先将所有节点的父节点缓存到字典中
    parents := map[int]*runtime.TreeNode{}
    var findParent func(root *runtime.TreeNode)
    findParent = func(root *runtime.TreeNode) {
        if root == nil {
            return
        }
        if root.Left != nil {
            parents[root.Left.Val] = root
        }
        if root.Right != nil {
            parents[root.Right.Val] = root
        }

        findParent(root.Left)
        findParent(root.Right)
    }
    findParent(root)

    // 从目标节点向三个方向寻找节点，分别是左、右和上
    // 每次找到一个节点k减1，到k=0时插入结果集并返回
    // 为了避免重复访问节点，每次寻找带上前驱节点
    res := make([]int, 0)
    var dfs func(root, from *runtime.TreeNode, k int)
    dfs = func(root, from *runtime.TreeNode, k int) {
        if k == 0 {
            res = append(res, root.Val)
            return
        }
        if root.Left != nil && root.Left != from {
            dfs(root.Left, root, k-1)
        }
        if root.Right != nil && root.Right != from {
            dfs(root.Right, root, k-1)
        }
        if parents[root.Val] != nil && parents[root.Val] != from {
            dfs(parents[root.Val], root, k-1)
        }
    }
    dfs(target, nil, k)

    return res
}
