package _0199

import (
    "github.com/raw34/leetcode/runtime"
    "math"
)

func rightSideView(root *runtime.TreeNode) []int {
    res := make([]int, 0)
    if root == nil {
        return res
    }
    queue := []*runtime.TreeNode{root}
    for len(queue) > 0 {
        n := len(queue)
        for i := 0; i < n; i++ {
            node := queue[0]
            queue = queue[1:]
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
            if i == n-1 {
                res = append(res, node.Val)
            }
        }
    }
    return res
}

func rightSideView2(root *runtime.TreeNode) []int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }
    maxLevel := math.MinInt32
    hash := map[int][]int{}
    var dfs func(root *runtime.TreeNode, level int)
    dfs = func(root *runtime.TreeNode, level int) {
        if root == nil {
            maxLevel = max(maxLevel, level)
            return
        }
        dfs(root.Left, level+1)
        dfs(root.Right, level+1)
        hash[level] = append(hash[level], root.Val)
    }
    dfs(root, 0)

    res := make([]int, 0)
    for i := 0; i < maxLevel; i++ {
        vals := hash[i]
        res = append(res, vals[len(vals)-1])
    }
    return res
}
