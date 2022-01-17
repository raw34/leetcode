package _0025

import "github.com/raw34/leetcode/runtime"

func reverseKGroup(head *runtime.ListNode, k int) *runtime.ListNode {
    reverse := func(head *runtime.ListNode) *runtime.ListNode {
        var prev *runtime.ListNode
        curr := head
        for curr != nil {
            next := curr.Next
            curr.Next = prev
            prev = curr
            curr = next
        }
        return prev
    }

    dummy := &runtime.ListNode{Next: head}
    pre, end := dummy, dummy
    for end.Next != nil {
        for i := 0; i < k && end != nil; i++ {
            end = end.Next
        }
        if end == nil {
            break
        }
        start := pre.Next
        next := end.Next
        end.Next = nil
        pre.Next = reverse(start)
        start.Next = next
        pre = start
        end = pre
    }

    return dummy.Next
}
