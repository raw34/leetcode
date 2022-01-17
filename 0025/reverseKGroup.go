package _0025

import "github.com/raw34/leetcode/runtime"

func reverseKGroup(head *runtime.ListNode, k int) *runtime.ListNode {
    reverse := func(head *runtime.ListNode, left, right int) *runtime.ListNode {
        dummy := &runtime.ListNode{Next: head}

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

    // 求链表长度
    curr := head
    n := 0
    for ; curr != nil; n++ {
        curr = curr.Next
    }
    // 循环调用反转链表II的方法
    curr = head
    for l, r := 1, k; r < n+1; r++ {
        if r%k == 0 {
            head = reverse(head, l, r)
            l = r + 1
        }
    }

    return head
}
