package _0103

import "github.com/raw34/leetcode/runtime"

func zigzagLevelOrder(root *runtime.TreeNode) [][]int {
    res := make([][]int, 0)
    if root == nil {
        return res
    }

    queue := []*runtime.TreeNode{root}
    for i := 0; len(queue) > 0; i++ {
        length := len(queue)
        curr := make([]int, 0)
        for j := 0; j < length; j++ {
            node := queue[0]
            queue = queue[1:]
            curr = append(curr, node.Val)

            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        if i%2 == 0 {
            res = append(res, curr)
        } else {
            res = append(res, []int{})
            for k := len(curr) - 1; k >= 0; k-- {
                res[i] = append(res[i], curr[k])
            }
        }
    }

    return res
}
