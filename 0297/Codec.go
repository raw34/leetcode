package _0297

import (
    "fmt"
    "github.com/raw34/leetcode/runtime"
    "strconv"
    "strings"
)

type Codec struct {
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *runtime.TreeNode) string {
    res := ""
    var dfs func(root *runtime.TreeNode)
    dfs = func(root *runtime.TreeNode) {
        if root == nil {
            res += "null,"
            return
        }
        res += fmt.Sprintf("%d,", root.Val)
        dfs(root.Left)
        dfs(root.Right)
    }
    dfs(root)

    return res
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *runtime.TreeNode {
    values := strings.Split(data, ",")
    var build func() *runtime.TreeNode
    build = func() *runtime.TreeNode {
        str := values[0]
        values = values[1:]
        if str == "null" {
            return nil
        }
        val, _ := strconv.Atoi(str)
        root := &runtime.TreeNode{Val: val}
        root.Left = build()
        root.Right = build()

        return root
    }

    return build()
}
