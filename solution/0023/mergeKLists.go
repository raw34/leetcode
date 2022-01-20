package _0023

import "github.com/raw34/leetcode/runtime"

func mergeKLists(lists []*runtime.ListNode) *runtime.ListNode {
    if len(lists) == 0 {
        return nil
    }

    // 转化为两个链表合并问题
    l1 := lists[0]
    lists = lists[1:]
    for len(lists) > 0 {
        l2 := lists[0]
        lists = lists[1:]
        l1 = mergeTwoLists(l1, l2)
    }

    return l1
}

func mergeTwoLists(l1 *runtime.ListNode, l2 *runtime.ListNode) *runtime.ListNode {
    dummy := &runtime.ListNode{}
    curr := dummy
    for l1 != nil && l2 != nil {
        v1, v2 := l1.Val, l2.Val
        if v1 < v2 {
            curr.Next = l1
            l1 = l1.Next
        } else {
            curr.Next = l2
            l2 = l2.Next
        }
        curr = curr.Next
    }

    if l1 != nil {
        curr.Next = l1
    }
    if l2 != nil {
        curr.Next = l2
    }

    return dummy.Next
}
