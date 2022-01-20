package _0143

import "github.com/raw34/leetcode/runtime"

func reorderList(head *runtime.ListNode) {
    if head == nil {
        return
    }

    nodes := []*runtime.ListNode{}
    node := head
    for node != nil {
        nodes = append(nodes, node)
        node = node.Next
    }

    i := 0
    j := len(nodes) - 1

    for i < j {
        nodes[i].Next = nodes[j]
        i++
        if i == j {
            break
        }
        nodes[j].Next = nodes[i]
        j--
    }
    nodes[j].Next = nil
}
