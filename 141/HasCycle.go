package _141

import "github.com/raw34/leetcode/runtime"

func hasCycle(head *runtime.ListNode) bool {
	if head == nil {
		return false
	}

	cache := make(map[*runtime.ListNode]int, 0)

	for head != nil {
		if cache[head] > 1 {
			return true
		}
		cache[head]++
		head = head.Next
	}

	return false
}