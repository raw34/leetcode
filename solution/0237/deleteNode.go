package _0237

import "github.com/raw34/leetcode/runtime"

func deleteNode(node *runtime.ListNode) {
    node.Val = node.Next.Val
    node.Next = node.Next.Next
}
