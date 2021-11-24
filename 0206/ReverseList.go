package _206

import "github.com/raw34/leetcode/runtime"

func reverseList(head *runtime.ListNode) *runtime.ListNode {
	var prev *runtime.ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}