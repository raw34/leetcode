package runtime

import (
    "testing"
)

func TestBinaryTree_build(t *testing.T) {
    nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    tree := &BinaryTree{}
    tree.build(nums)
}

func TestBinaryTree_buildFromPreorderAndInorder1(t *testing.T) {
    preorder := []int{3, 9, 20, 15, 7}
    inorder := []int{9, 3, 15, 20, 7}
    tree := &BinaryTree{}
    root := tree.buildFromPreorderAndInorder1(preorder, inorder)
    tree.displayPreorder1(root)
    tree.displayInorder1(root)
    tree.displayPostorder1(root)
}

func TestBinaryTree_buildFromPreorderAndInorder2(t *testing.T) {
    preorder := []int{3, 9, 20, 15, 7}
    inorder := []int{9, 3, 15, 20, 7}
    tree := &BinaryTree{}
    root := tree.buildFromPreorderAndInorder2(preorder, inorder)
    tree.displayPreorder1(root)
    tree.displayInorder1(root)
    tree.displayPostorder1(root)
}
