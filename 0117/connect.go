package _0117

import "github.com/raw34/leetcode/runtime"

func connect(root *runtime.Node) *runtime.Node {
    if root == nil {
        return nil
    }
    queue := []*runtime.Node{root}
    for len(queue) > 0 {
        length := len(queue)
        for i := 0; i < length; i++ {
            node := queue[0]
            queue = queue[1:]

            if i < length-1 {
                node.Next = queue[0]
            }

            if node.Left != nil {
                queue = append(queue, node.Left)
            }

            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
    }

    return root
}
