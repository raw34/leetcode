package _1028

import "github.com/raw34/leetcode/runtime"

func recoverFromPreorder(traversal string) *runtime.TreeNode {
    nodes := map[int]*runtime.TreeNode{-1: &runtime.TreeNode{}}
    addNode := func(dep, val int) {
        nodes[dep] = &runtime.TreeNode{Val: val}
        if nodes[dep-1].Left == nil {
            nodes[dep-1].Left = nodes[dep]
        } else {
            nodes[dep-1].Right = nodes[dep]
        }
    }

    dep := 0
    val := 0
    for i := 0; i < len(traversal); i++ {
        curr := traversal[i]
        if curr == '-' && val > 0 {
            addNode(dep, val)
            dep = 0
            val = 0
        }
        if curr != '-' {
            val = val*10 + int(curr-'0')
        } else {
            dep++
        }
    }
    addNode(dep, val)

    return nodes[0]
}