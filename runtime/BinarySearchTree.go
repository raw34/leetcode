package runtime

type BinarySearchTree struct {
}

func NewBinarySearchTree() BinarySearchTree {
    return BinarySearchTree{}
}

func (this *BinarySearchTree) BuildFromPreorder(preorder []int) *TreeNode {
    var dfs func(left, right int) *TreeNode
    dfs = func(left, right int) *TreeNode {
        if left > right {
            return nil
        }

        val := preorder[0]
        root := &TreeNode{Val: val}

        // 寻找左右子树中间节点
        mid := right + 1
        for i := 0; i < len(preorder); i++ {
            if preorder[i] > val {
                mid = i
                break
            }
        }
        root.Left = this.BuildFromPreorder(preorder[left+1 : mid])
        root.Right = this.BuildFromPreorder(preorder[mid:])

        return root
    }

    return dfs(0, len(preorder)-1)
}
