package runtime

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestBinaryTree_build(t *testing.T) {
    str := "1-2--3"
    for i := 0; i < len(str); i++ {
        fmt.Println(str[0:i])
    }

    nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    tree := &BinaryTree{}
    root := tree.Build(nums)
    res1 := tree.displayPreorder1(root)
    fmt.Println(res1)
    res2 := tree.displayInorder1(root)
    fmt.Println(res2)
    res3 := tree.displayPostorder1(root)
    fmt.Println(res3)
    res4 := tree.DisplayLevelOrder1(root)
    fmt.Println(res4)
}

func TestBinaryTree_buildFromPreorderAndInorder1(t *testing.T) {
    preorder := []int{1, 2, 4, 8, 9, 5, 10, 11, 3, 6, 12, 13, 7, 14, 15}
    inorder := []int{8, 4, 9, 2, 10, 5, 11, 1, 12, 6, 13, 3, 14, 7, 15}
    tree := &BinaryTree{}
    root := tree.buildFromPreorderAndInorder1(preorder, inorder)
    res1 := tree.displayPreorder1(root)
    assert.Equal(t, []int{1, 2, 4, 8, 9, 5, 10, 11, 3, 6, 12, 13, 7, 14, 15}, res1)
    res2 := tree.displayInorder1(root)
    assert.Equal(t, []int{8, 4, 9, 2, 10, 5, 11, 1, 12, 6, 13, 3, 14, 7, 15}, res2)
    res3 := tree.displayPostorder1(root)
    assert.Equal(t, []int{8, 9, 4, 10, 11, 5, 2, 12, 13, 6, 14, 15, 7, 3, 1}, res3)
    res4 := tree.DisplayLevelOrder1(root)
    assert.Equal(t, [][]int{{1}, {2, 3}, {4, 5, 6, 7}, {8, 9, 10, 11, 12, 13, 14, 15}}, res4)
}

func TestBinaryTree_buildFromPreorderAndInorder2(t *testing.T) {
    preorder := []int{1, 2, 4, 8, 9, 5, 10, 11, 3, 6, 12, 13, 7, 14, 15}
    inorder := []int{8, 4, 9, 2, 10, 5, 11, 1, 12, 6, 13, 3, 14, 7, 15}
    tree := &BinaryTree{}
    root := tree.buildFromPreorderAndInorder2(preorder, inorder)
    res1 := tree.displayPreorder2(root)
    assert.Equal(t, []int{1, 2, 4, 8, 9, 5, 10, 11, 3, 6, 12, 13, 7, 14, 15}, res1)
    res2 := tree.displayInorder2(root)
    assert.Equal(t, []int{8, 4, 9, 2, 10, 5, 11, 1, 12, 6, 13, 3, 14, 7, 15}, res2)
    res3 := tree.displayPostorder2(root)
    assert.Equal(t, []int{8, 9, 4, 10, 11, 5, 2, 12, 13, 6, 14, 15, 7, 3, 1}, res3)
    res4 := tree.DisplayLevelOrder2(root)
    assert.Equal(t, [][]int{{1}, {2, 3}, {4, 5, 6, 7}, {8, 9, 10, 11, 12, 13, 14, 15}}, res4)
}

func TestBinaryTree_buildFromInorderAndPostorder1(t *testing.T) {
    inorder := []int{8, 4, 9, 2, 10, 5, 11, 1, 12, 6, 13, 3, 14, 7, 15}
    postorder := []int{8, 9, 4, 10, 11, 5, 2, 12, 13, 6, 14, 15, 7, 3, 1}
    tree := &BinaryTree{}
    root := tree.buildFromInorderAndPostorder1(inorder, postorder)
    res1 := tree.displayPreorder1(root)
    assert.Equal(t, []int{1, 2, 4, 8, 9, 5, 10, 11, 3, 6, 12, 13, 7, 14, 15}, res1)
    res2 := tree.displayInorder1(root)
    assert.Equal(t, []int{8, 4, 9, 2, 10, 5, 11, 1, 12, 6, 13, 3, 14, 7, 15}, res2)
    res3 := tree.displayPostorder1(root)
    assert.Equal(t, []int{8, 9, 4, 10, 11, 5, 2, 12, 13, 6, 14, 15, 7, 3, 1}, res3)
    res4 := tree.DisplayLevelOrder1(root)
    assert.Equal(t, [][]int{{1}, {2, 3}, {4, 5, 6, 7}, {8, 9, 10, 11, 12, 13, 14, 15}}, res4)
}

func TestBinaryTree_buildFromInorderAndPostorder2(t *testing.T) {
    inorder := []int{8, 4, 9, 2, 10, 5, 11, 1, 12, 6, 13, 3, 14, 7, 15}
    postorder := []int{8, 9, 4, 10, 11, 5, 2, 12, 13, 6, 14, 15, 7, 3, 1}
    tree := &BinaryTree{}
    root := tree.buildFromInorderAndPostorder2(inorder, postorder)
    res1 := tree.displayPreorder2(root)
    assert.Equal(t, []int{1, 2, 4, 8, 9, 5, 10, 11, 3, 6, 12, 13, 7, 14, 15}, res1)
    res2 := tree.displayInorder2(root)
    assert.Equal(t, []int{8, 4, 9, 2, 10, 5, 11, 1, 12, 6, 13, 3, 14, 7, 15}, res2)
    res3 := tree.displayPostorder2(root)
    assert.Equal(t, []int{8, 9, 4, 10, 11, 5, 2, 12, 13, 6, 14, 15, 7, 3, 1}, res3)
    res4 := tree.DisplayLevelOrder2(root)
    assert.Equal(t, [][]int{{1}, {2, 3}, {4, 5, 6, 7}, {8, 9, 10, 11, 12, 13, 14, 15}}, res4)
}

func TestBinaryTree_buildFromPreorderAndPostorder1(t *testing.T) {
    preorder := []int{1, 2, 4, 8, 9, 5, 10, 11, 3, 6, 12, 13, 7, 14, 15}
    postorder := []int{8, 9, 4, 10, 11, 5, 2, 12, 13, 6, 14, 15, 7, 3, 1}
    tree := &BinaryTree{}
    root := tree.buildFromPreorderAndPostorder1(preorder, postorder)
    res1 := tree.displayPreorder1(root)
    assert.Equal(t, []int{1, 2, 4, 8, 9, 5, 10, 11, 3, 6, 12, 13, 7, 14, 15}, res1)
    res2 := tree.displayInorder1(root)
    assert.Equal(t, []int{8, 4, 9, 2, 10, 5, 11, 1, 12, 6, 13, 3, 14, 7, 15}, res2)
    res3 := tree.displayPostorder1(root)
    assert.Equal(t, []int{8, 9, 4, 10, 11, 5, 2, 12, 13, 6, 14, 15, 7, 3, 1}, res3)
    res4 := tree.DisplayLevelOrder1(root)
    assert.Equal(t, [][]int{{1}, {2, 3}, {4, 5, 6, 7}, {8, 9, 10, 11, 12, 13, 14, 15}}, res4)
}

func TestBinaryTree_buildFromPreorderAndPostorder2(t *testing.T) {
    preorder := []int{1, 2, 4, 8, 9, 5, 10, 11, 3, 6, 12, 13, 7, 14, 15}
    postorder := []int{8, 9, 4, 10, 11, 5, 2, 12, 13, 6, 14, 15, 7, 3, 1}
    tree := &BinaryTree{}
    root := tree.buildFromPreorderAndPostorder2(preorder, postorder)
    res1 := tree.displayPreorder2(root)
    assert.Equal(t, []int{1, 2, 4, 8, 9, 5, 10, 11, 3, 6, 12, 13, 7, 14, 15}, res1)
    res2 := tree.displayInorder2(root)
    assert.Equal(t, []int{8, 4, 9, 2, 10, 5, 11, 1, 12, 6, 13, 3, 14, 7, 15}, res2)
    res3 := tree.displayPostorder2(root)
    assert.Equal(t, []int{8, 9, 4, 10, 11, 5, 2, 12, 13, 6, 14, 15, 7, 3, 1}, res3)
    res4 := tree.DisplayLevelOrder2(root)
    assert.Equal(t, [][]int{{1}, {2, 3}, {4, 5, 6, 7}, {8, 9, 10, 11, 12, 13, 14, 15}}, res4)
}
