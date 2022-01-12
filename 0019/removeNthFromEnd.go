package _0019

import "github.com/raw34/leetcode/runtime"

func removeNthFromEnd(head *runtime.ListNode, n int) *runtime.ListNode {
    // 将链表存到map里，以便利用索引寻找节点
    hash := map[int]*runtime.ListNode{}
    dummy := &runtime.ListNode{Next: head}
    curr := dummy
    i := 0
    for curr != nil {
        hash[i] = curr
        curr = curr.Next
        i++
    }

    // 从map中寻找并删除节点
    hash[i-n-1].Next = hash[i-n+1]

    return dummy.Next
}
