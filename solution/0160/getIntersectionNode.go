package _0160

import "github.com/raw34/leetcode/runtime"

func getIntersectionNode(headA, headB *runtime.ListNode) *runtime.ListNode {
    hash := map[*runtime.ListNode]*runtime.ListNode{}
    for headA != nil {
        hash[headA] = headA
        headA = headA.Next
    }
    for headB != nil {
        if head, ok := hash[headB]; ok {
            return head
        }
        headB = headB.Next
    }
    return nil
}
