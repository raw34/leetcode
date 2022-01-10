package _0148

import "github.com/raw34/leetcode/runtime"

func sortList(head *runtime.ListNode) *runtime.ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    slow := head
    fast := head.Next
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    tail := slow.Next
    slow.Next = nil

    return merge(sortList(head), sortList(tail))
}

func merge(head1, head2 *runtime.ListNode) *runtime.ListNode {
    dummy := &runtime.ListNode{}
    curr := dummy

    for head1 != nil && head2 != nil {
        if head1.Val < head2.Val {
            curr.Next = head1
            head1 = head1.Next
        } else {
            curr.Next = head2
            head2 = head2.Next
        }
        curr = curr.Next
    }

    if head1 != nil {
        curr.Next = head1
    } else {
        curr.Next = head2
    }

    return dummy.Next
}
