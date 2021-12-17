package runtime

type Node struct {
    Val    int
    Parent *Node
    Left   *Node
    Right  *Node
    Next   *Node
}

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}
