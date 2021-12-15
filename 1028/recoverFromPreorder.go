package _1028

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func recoverFromPreorder(traversal string) *TreeNode {
    nodes := map[int]*TreeNode{-1: &TreeNode{}}
    addNode := func(dep, val int) {
        nodes[dep] = &TreeNode{Val: val}
        if nodes[dep-1].Left == nil {
            nodes[dep-1].Left = nodes[dep]
        } else {
            nodes[dep-1].Right = nodes[dep]
        }
    }

    dep := 0
    val := 0
    hasVal := false
    for i := 0; i < len(traversal); i++ {
        curr := traversal[i]
        if curr != '-' {
            val = val*10 + int(curr-'0')
            hasVal = true
        } else {
            if hasVal {
                addNode(dep, val)
                dep = 0
                val = 0
                hasVal = false
            }
            dep++
        }
    }
    addNode(dep, val)

    return nodes[0]
}
