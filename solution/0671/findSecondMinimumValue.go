package _0671

import (
    "github.com/raw34/leetcode/runtime"
    "math"
)

func findSecondMinimumValue(root *runtime.TreeNode) int {
    stack := [2]int{math.MaxInt64, math.MaxInt64}
    queue := []*runtime.TreeNode{root}
    for len(queue) > 0 {
        root = queue[0]
        queue = queue[1:]

        if root.Val > stack[0] && root.Val < stack[1] {
            stack[1] = root.Val
        } else if root.Val < stack[0] {
            stack[0], stack[1] = root.Val, stack[0]
        }

        if root.Left != nil {
            queue = append(queue, root.Left)
        }
        if root.Right != nil {
            queue = append(queue, root.Right)
        }
    }

    if stack[1] == math.MaxInt64 {
        return -1
    }

    return stack[1]
}
