package _0002

import "github.com/raw34/leetcode/runtime"

func addTwoNumbers(l1 *runtime.ListNode, l2 *runtime.ListNode) *runtime.ListNode {
    head := &runtime.ListNode{}
    curr := head
    carry := 0
    for l1 != nil || l2 != nil {
        val1, val2 := 0, 0
        if l1 != nil {
            val1 = l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            val2 = l2.Val
            l2 = l2.Next
        }
        sum := val1 + val2 + carry
        carry = sum / 10
        sum = sum % 10
        curr.Next = &runtime.ListNode{Val: sum}
        curr = curr.Next
    }

    if carry > 0 {
        curr.Next = &runtime.ListNode{Val: carry}
    }

    return head.Next
}
