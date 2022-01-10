package _0092

import "github.com/raw34/leetcode/runtime"

func reverseBetween(head *runtime.ListNode, left int, right int) *runtime.ListNode {
    dummy := &runtime.ListNode{Val: -1, Next: head}

    prev := dummy
    for i := 0; i < left-1; i++ {
        prev = prev.Next
    }

    curr := prev.Next
    for i := 0; i < right-left; i++ {
        next := curr.Next
        curr.Next = next.Next
        next.Next = prev.Next
        prev.Next = next
    }

    return dummy.Next
}
