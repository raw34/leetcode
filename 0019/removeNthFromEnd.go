package _0019

import "github.com/raw34/leetcode/runtime"

func removeNthFromEnd(head *runtime.ListNode, n int) *runtime.ListNode {
    hash := map[int]*runtime.ListNode{}
    dummy := &runtime.ListNode{Next: head}
    curr := dummy
    i := 0
    for curr != nil {
        hash[i] = curr
        curr = curr.Next
        i++
    }

    hash[i-n-1].Next = hash[i-n+1]

    return dummy.Next
}
