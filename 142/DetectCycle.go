package _142

import "github.com/raw34/leetcode/runtime"

func detectCycle(head *runtime.ListNode) *runtime.ListNode {
	slow := head
	fast := head
	for true {
		if fast == nil || fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}

	fast = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return fast
}