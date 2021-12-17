package _0510

import "github.com/raw34/leetcode/runtime"

func inorderSuccessor(node *runtime.Node) *runtime.Node {
    if node.Right != nil {
        node = node.Right
        for node.Left != nil {
            node = node.Left
        }
        return node
    }

    for node.Parent != nil && node.Parent.Right == node {
        node = node.Parent
    }

    return node.Parent
}
