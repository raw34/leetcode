package _0109

import "github.com/raw34/leetcode/runtime"

func sortedListToBST(head *runtime.ListNode) *runtime.TreeNode {
    curr := head
    length := 0
    for curr != nil {
        curr = curr.Next
        length++
    }

    var dfs func(left, right int) *runtime.TreeNode
    dfs = func(left, right int) *runtime.TreeNode {
        if left > right {
            return nil
        }
        mid := (left + right + 1) / 2
        root := &runtime.TreeNode{}
        root.Left = dfs(left, mid-1)
        root.Val = head.Val
        head = head.Next
        root.Right = dfs(mid+1, right)

        return root
    }

    return dfs(0, length-1)
}
