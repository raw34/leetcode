package runtime

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestBinaryTree_build(t *testing.T) {
    nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    tree := &BinaryTree{}
    root := tree.build(nums)
    res1 := tree.displayPreorder1(root)
    fmt.Println(res1)
    res2 := tree.displayInorder1(root)
    fmt.Println(res2)
    res3 := tree.displayPostorder1(root)
    fmt.Println(res3)
    res4 := tree.displayLevelOrder1(root)
    fmt.Println(res4)
}

func TestBinaryTree_buildFromPreorderAndInorder1(t *testing.T) {
    preorder := []int{3, 9, 20, 15, 7}
    inorder := []int{9, 3, 15, 20, 7}
    tree := &BinaryTree{}
    root := tree.buildFromPreorderAndInorder1(preorder, inorder)
    res1 := tree.displayPreorder1(root)
    assert.Equal(t, []int{3, 9, 20, 15, 7}, res1)
    res2 := tree.displayInorder1(root)
    assert.Equal(t, []int{9, 3, 15, 20, 7}, res2)
    res3 := tree.displayPostorder1(root)
    assert.Equal(t, []int{9, 15, 7, 20, 3}, res3)
    res4 := tree.displayLevelOrder1(root)
    assert.Equal(t, [][]int{{3}, {9, 20}, {15, 7}}, res4)
}

func TestBinaryTree_buildFromPreorderAndInorder2(t *testing.T) {
    preorder := []int{3, 9, 20, 15, 7}
    inorder := []int{9, 3, 15, 20, 7}
    tree := &BinaryTree{}
    root := tree.buildFromPreorderAndInorder2(preorder, inorder)
    res1 := tree.displayPreorder1(root)
    assert.Equal(t, []int{3, 9, 20, 15, 7}, res1)
    res2 := tree.displayInorder1(root)
    assert.Equal(t, []int{9, 3, 15, 20, 7}, res2)
    res3 := tree.displayPostorder1(root)
    assert.Equal(t, []int{9, 15, 7, 20, 3}, res3)
    res4 := tree.displayLevelOrder1(root)
    assert.Equal(t, [][]int{{3}, {9, 20}, {15, 7}}, res4)
}
