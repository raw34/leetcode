package _0147

import "github.com/raw34/leetcode/runtime"

func insertionSortList(head *runtime.ListNode) *runtime.ListNode {
    if head == nil {
        return head
    }
    dummy := &runtime.ListNode{Val: -1, Next: head}
    last := head
    curr := head.Next
    for curr != nil {
        if last.Val <= curr.Val {
            last = last.Next
        } else {
            prev := dummy
            for prev.Next.Val <= curr.Val {
                prev = prev.Next
            }
            last.Next = curr.Next
            curr.Next = prev.Next
            prev.Next = curr
        }
        curr = last.Next
    }

    return dummy.Next
}
