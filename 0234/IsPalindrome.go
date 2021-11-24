package _234

import "github.com/raw34/leetcode/runtime"

func isPalindrome(head *runtime.ListNode) bool {
	if head == nil {
		return false
	}

	if head.Next == nil {
		return true
	}

	stack1, stack2 := make([]int, 0), make([]int, 0)
	slow, fast := head, head.Next
	for fast != nil {
		stack1 = append(stack1, slow.Val)
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	for slow != nil {
		stack2 = append(stack2, slow.Val)
		slow = slow.Next
	}

	len1 := len(stack1)
	len2 := len(stack2)
	if len1 < len2 {
		stack2 = stack2[1:]
	}

	for i := 0; i < len1; i++ {
		j := len1 - i - 1
		if stack1[i] != stack2[j] {
			return false
		}
	}

	return true
}