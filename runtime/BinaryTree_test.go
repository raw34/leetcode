package runtime

import (
    "fmt"
    "testing"
)

func TestBinaryTree_build(t *testing.T) {
    nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    tree := &BinaryTree{}
    tree.build(nums)

    fmt.Println(tree)
}

func TestBinaryTree_displayInorder(t *testing.T) {
}

func TestBinaryTree_displayPostorder(t *testing.T) {
}

func TestBinaryTree_displayPreorder(t *testing.T) {
}
