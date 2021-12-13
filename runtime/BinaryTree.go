package runtime

import "fmt"

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

func (bt *BinaryTree) buildFromPreorderAndInorder1(preorder []int, inorder []int) *TreeNode {
    indexMap := make(map[int]int, 0)
    for i, v := range inorder {
        indexMap[v] = i
    }

    var dfs func(preorderLeft, preorderRight, inorderLeft, inorderRight int) *TreeNode
    dfs = func(preorderLeft, preorderRight, inorderLeft, inorderRight int) *TreeNode {
        if preorderLeft > preorderRight {
            return nil
        }
        preorderRoot := preorderLeft
        val := preorder[preorderRoot]
        inorderRoot := indexMap[val]
        leftTreeSize := inorderRoot - inorderLeft

        root := &TreeNode{Val: val}
        root.Left = dfs(preorderLeft+1, preorderLeft+leftTreeSize, inorderLeft, inorderRoot-1)
        root.Right = dfs(preorderLeft+leftTreeSize+1, preorderRight, inorderRoot+1, inorderRight)

        return root
    }

    n := len(preorder)
    return dfs(0, n-1, 0, n-1)
}

func (bt *BinaryTree) buildFromPreorderAndInorder2(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    root := &TreeNode{Val: preorder[0]}
    stack := []*TreeNode{root}
    inorderIndex := 0
    for i := 1; i < len(preorder); i++ {
        preorderVal := preorder[i]
        node := stack[len(stack)-1]
        if node.Val != inorder[inorderIndex] {
            node.Left = &TreeNode{Val: preorderVal}
            stack = append(stack, node.Left)
        } else {
            for len(stack) != 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
                node = stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                inorderIndex++
            }
            node.Right = &TreeNode{Val: preorderVal}
            stack = append(stack, node.Right)
        }
    }
    return root
}

func (bt *BinaryTree) displayPreorder1(root *TreeNode) []int {
    stack := make([]int, 0)

    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return
        }
        stack = append(stack, root.Val)
        dfs(root.Left)
        dfs(root.Right)
    }
    dfs(root)

    fmt.Println(stack)
    return stack
}

func (bt *BinaryTree) displayInorder1(root *TreeNode) []int {
    stack := make([]int, 0)

    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return
        }
        dfs(root.Left)
        stack = append(stack, root.Val)
        dfs(root.Right)
    }
    dfs(root)

    fmt.Println(stack)
    return stack
}

func (bt *BinaryTree) displayPostorder1(root *TreeNode) []int {
    stack := make([]int, 0)

    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return
        }
        dfs(root.Left)
        dfs(root.Right)
        stack = append(stack, root.Val)
    }
    dfs(root)

    fmt.Println(stack)
    return stack
}
