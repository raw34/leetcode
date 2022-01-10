package _0083

import "github.com/raw34/leetcode/runtime"

func deleteDuplicates(head *runtime.ListNode) *runtime.ListNode {
    if head == nil {
        return head
    }

    curr := head
    for curr.Next != nil {
        if curr.Val == curr.Next.Val {
            curr.Next = curr.Next.Next
        } else {
            curr = curr.Next
        }
    }

    return head
}
