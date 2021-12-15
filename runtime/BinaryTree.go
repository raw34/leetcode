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
    root := &TreeNode{Val: nums[0]}
    queue := []*TreeNode{root}

    n := len(nums)
    i := 1
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]

        if i < n {
            node.Left = &TreeNode{Val: nums[i]}
            queue = append(queue, node.Left)
        }
        i++

        if i < n {
            node.Right = &TreeNode{Val: nums[i]}
            queue = append(queue, node.Right)
        }
        i++
    }

    return root
}

func (bt *BinaryTree) buildFromPreorderAndInorder1(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 || len(inorder) == 0 {
        return nil
    }

    indexMap := make(map[int]int, 0)
    for i, v := range inorder {
        indexMap[v] = i
    }

    val := preorder[0]
    mid := indexMap[val]
    root := &TreeNode{Val: val}
    root.Left = bt.buildFromPreorderAndInorder1(preorder[1:mid+1], inorder[:mid])
    root.Right = bt.buildFromPreorderAndInorder1(preorder[mid+1:], inorder[mid+1:])

    return root
}

func (bt *BinaryTree) buildFromPreorderAndInorder2(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 || len(inorder) == 0 {
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

func (bt *BinaryTree) buildFromInorderAndPostorder1(inorder []int, postorder []int) *TreeNode {
    if len(inorder) == 0 || len(postorder) == 0 {
        return nil
    }

    indexMap := make(map[int]int, 0)
    for i, v := range inorder {
        indexMap[v] = i
    }

    val := postorder[len(postorder)-1]
    mid := indexMap[val]
    root := &TreeNode{Val: val}
    root.Left = bt.buildFromInorderAndPostorder1(inorder[:mid], postorder[:mid])
    root.Right = bt.buildFromInorderAndPostorder1(inorder[mid+1:], postorder[mid:len(postorder)-1])

    return root
}

func (bt *BinaryTree) buildFromInorderAndPostorder2(inorder []int, postorder []int) *TreeNode {
    return nil
}

func (bt *BinaryTree) buildFromPreorderAndPostorder1(preorder []int, postorder []int) *TreeNode {
    if len(preorder) == 0 || len(postorder) == 0 {
        return nil
    }

    val := preorder[0]
    root := &TreeNode{Val: val}
    if len(postorder) == 1 {
        return root
    }

    mid := 0
    for i := 0; i < len(postorder); i++ {
        if postorder[i] == preorder[1] {
            mid = i + 1
            break
        }
    }

    root.Left = bt.buildFromPreorderAndPostorder1(preorder[1:mid+1], postorder[0:mid])
    root.Right = bt.buildFromPreorderAndPostorder1(preorder[mid+1:], postorder[mid:len(postorder)-1])

    return root
}

func (bt *BinaryTree) buildFromPreorderAndPostorder2(preorder []int, postorder []int) *TreeNode {
    return nil
}

/*
前序遍历顺序为：根 -> 左 -> 右
*/
func (bt *BinaryTree) displayPreorder1(root *TreeNode) []int {
    res := make([]int, 0)

    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return
        }
        res = append(res, root.Val)
        dfs(root.Left)
        dfs(root.Right)
    }
    dfs(root)

    return res
}

/*
前序遍历顺序为：根 -> 左 -> 右
解题思路：
1、从当前点开始，向左遍历，直到最左叶子节点，遍历过程中，取出当前节点值，并缓存当前节点到临时栈
2、如果临时栈不为空，出栈一个节点，将该节点的右叶子节点赋值给当前节点
3、当节点不为空或临时栈不为空时，循环1、2步
*/
func (bt *BinaryTree) displayPreorder2(root *TreeNode) []int {
    res := make([]int, 0)
    stack := make([]*TreeNode, 0)
    for root != nil || len(stack) > 0 {
        for root != nil {
            res = append(res, root.Val)
            stack = append(stack, root)
            root = root.Left
        }
        if len(stack) > 0 {
            root = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            root = root.Right
        }
    }

    return res
}

/*
中序遍历顺序为：左 -> 根 -> 右
*/
func (bt *BinaryTree) displayInorder1(root *TreeNode) []int {
    res := make([]int, 0)

    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return
        }
        dfs(root.Left)
        res = append(res, root.Val)
        dfs(root.Right)
    }
    dfs(root)

    return res
}

/*
中序遍历顺序为：左 -> 根 -> 右
解题思路：
1、从当前点开始，向左遍历，直到最左叶子节点，遍历过程中，缓存当前节点到临时栈
2、如果临时栈不为空，出栈一个节点，取出节点值，将该节点的右叶子节点赋值给当前节点
3、当节点不为空或临时栈不为空时，循环1、2步
*/
func (bt *BinaryTree) displayInorder2(root *TreeNode) []int {
    res := make([]int, 0)
    stack := make([]*TreeNode, 0)
    for root != nil || len(stack) > 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }

        if len(stack) > 0 {
            root = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            res = append(res, root.Val)
            root = root.Right
        }
    }

    return res
}

/*
后序遍历顺序为：左 -> 右 -> 根
*/
func (bt *BinaryTree) displayPostorder1(root *TreeNode) []int {
    res := make([]int, 0)

    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return
        }
        dfs(root.Left)
        dfs(root.Right)
        res = append(res, root.Val)
    }
    dfs(root)

    return res
}

/*
后序遍历顺序为：左 -> 右 -> 根
解题思路：
1、从当前节点开始，向右遍历，直到最右叶子节点，遍历过程中，取出当前节点值，缓存当前节点到临时栈
2、如果当前节点为空，出栈一个节点，将节点的左叶子节点赋值给当前节点
3、当节点不为空或临时栈不为空时，循环1、2步
*/
func (bt *BinaryTree) displayPostorder2(root *TreeNode) []int {
    res := make([]int, 0)
    stack := make([]*TreeNode, 0)
    for root != nil || len(stack) > 0 {
        if root != nil {
            res = append([]int{root.Val}, res...)
            stack = append(stack, root)
            root = root.Right
        } else {
            root = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            root = root.Left
        }
    }

    return res
}

// DisplayLevelOrder1 /*
/*
层序遍历顺序为：每一层从左到右
*/
func (bt *BinaryTree) DisplayLevelOrder1(root *TreeNode) [][]int {
    res := make([][]int, 0)
    var dfs func(root *TreeNode, level int)
    dfs = func(root *TreeNode, level int) {
        if root == nil {
            return
        }
        if len(res) == level {
            res = append(res, []int{})
        }
        res[level] = append(res[level], root.Val)

        if root.Left != nil {
            dfs(root.Left, level+1)
        }

        if root.Right != nil {
            dfs(root.Right, level+1)
        }
    }
    dfs(root, 0)

    return res
}

// DisplayLevelOrder2 /*
/* 层序遍历顺序为：每一层从左到右
解题思路：
1、初始化临时队列，将根节点入队
2、遍历临时队列，每次出队一个节点，获取当前节点值，并将当前节点左右叶子节点依次压入临时队列
3、当临时队列不为空时，循环1、2步
*/
func (bt *BinaryTree) DisplayLevelOrder2(root *TreeNode) [][]int {
    res := make([][]int, 0)
    queue := []*TreeNode{root}
    for level := 0; len(queue) > 0; level++ {
        length := len(queue)
        res = append(res, []int{})
        for i := 0; i < length; i++ {
            node := queue[0]
            queue = queue[1:]
            res[level] = append(res[level], node.Val)

            if node.Left != nil {
                queue = append(queue, node.Left)
            }

            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        //fmt.Println(res[level])
    }

    return res
}
