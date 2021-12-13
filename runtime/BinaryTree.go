package runtime

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

type BinaryTree struct {
    stackPre  []int
    stackIn   []int
    stackPost []int
}

func (bt *BinaryTree) build(nums []int) *TreeNode {
    root := &TreeNode{nums[0], nil, nil}
    queue := []*TreeNode{root}

    n := len(nums)
    i := 1
    for len(queue) > 0 {
        node := queue[0]

        if i < n {
            node.Left = &TreeNode{nums[i], nil, nil}
            queue = append(queue, node.Left)
        }
        i++

        if i < n {
            node.Right = &TreeNode{nums[i], nil, nil}
            queue = append(queue, node.Right)
        }
        i++
    }

    return root
}

func (bt *BinaryTree) displayPreorder(root *TreeNode) []int {
    stack := make([]int, 0)

    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        stack = append(stack, root.Val)
        dfs(root.Left)
        dfs(root.Right)
    }
    dfs(root)

    return stack
}

func (bt *BinaryTree) displayInorder(root *TreeNode) []int {
    stack := make([]int, 0)

    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        dfs(root.Left)
        stack = append(stack, root.Val)
        dfs(root.Right)
    }
    dfs(root)

    return stack
}

func (bt *BinaryTree) displayPostorder(root *TreeNode) []int {
    stack := make([]int, 0)

    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        dfs(root.Left)
        dfs(root.Right)
        stack = append(stack, root.Val)
    }
    dfs(root)

    return stack
}
