package _21

import "github.com/raw34/leetcode/runtime"

func mergeTwoLists(l1 *runtime.ListNode, l2 *runtime.ListNode) *runtime.ListNode {
	dummy := &runtime.ListNode{}
	curr := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
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
	} else {
		curr.Next = l2
	}

	return dummy.Next
}