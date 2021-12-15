package _0449

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
    if root == nil {
        return "null"
    }

    return fmt.Sprintf("%d,%s,%s", root.Val, this.serialize(root.Left), this.serialize(root.Right))
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *runtime.TreeNode {
    values := strings.Split(data, ",")
    var dfs func() *runtime.TreeNode
    dfs = func() *runtime.TreeNode {
        str := values[0]
        values = values[1:]
        if str == "null" {
            return nil
        }
        val, _ := strconv.Atoi(str)
        root := &runtime.TreeNode{Val: val}
        root.Left = dfs()
        root.Right = dfs()
        return root
    }

    return dfs()
}
