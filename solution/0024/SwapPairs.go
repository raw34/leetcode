package _0024

import "github.com/raw34/leetcode/runtime"

func swapPairs(head *runtime.ListNode) *runtime.ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    newHead := head.Next
    head.Next = swapPairs(newHead.Next)
    newHead.Next = head

    return newHead
}
