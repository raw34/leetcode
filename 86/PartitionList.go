package _86

import "github.com/raw34/leetcode/runtime"

func partition(head *runtime.ListNode, x int) *runtime.ListNode {
	left := &runtime.ListNode{}
	right := &runtime.ListNode{}

	leftHead := left
	rightHead := right

	for head != nil {
		if head.Val < x {
			left.Next = head
			left= left.Next
		} else {
			right.Next = head
			right = right.Next
		}
		head = head.Next
	}

	right.Next = nil
	left.Next = rightHead.Next

	return leftHead.Next
}