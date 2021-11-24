package _82

import "github.com/raw34/leetcode/runtime"



func deleteDuplicates(head *runtime.ListNode) *runtime.ListNode {
	if head == nil {
		return head
	}

	dummy := &runtime.ListNode{Next: head}

	curr := dummy
	for curr.Next != nil && curr.Next.Next != nil {
		if curr.Next.Val == curr.Next.Next.Val {
			val := curr.Next.Val
			for curr.Next != nil && curr.Next.Val == val {
				curr.Next = curr.Next.Next
			}
		} else {
			curr =  curr.Next
		}
	}

	return dummy.Next
}