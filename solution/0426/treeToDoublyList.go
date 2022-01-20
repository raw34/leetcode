package _0426

import "github.com/raw34/leetcode/runtime"

func treeToDoublyList(root *runtime.Node) *runtime.Node {
    if root == nil {
        return nil
    }
    dummy := &runtime.Node{}
    prev := dummy
    var dfs func(node *runtime.Node)
    dfs = func(node *runtime.Node) {
        if node == nil {
            return
        }
        dfs(node.Left)
        node.Left = prev
        prev.Right = node
        prev = node
        dfs(node.Right)
    }
    dfs(root)

    head := dummy.Right
    head.Left, prev.Right = prev, head

    return head
}

func treeToDoublyList2(root *runtime.Node) *runtime.Node {
    if root == nil {
        return nil
    }
    dummy := &runtime.Node{}
    prev, curr := dummy, root
    stack := make([]*runtime.Node, 0)
    for curr != nil || len(stack) > 0 {
        for curr != nil {
            stack = append(stack, curr)
            curr = curr.Left
        }
        if len(stack) > 0 {
            curr = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            prev.Right = curr
            curr.Left = prev
            prev = curr
            curr = curr.Right
        }
    }
    head := dummy.Right
    head.Left, prev.Right = prev, head

    return head
}
