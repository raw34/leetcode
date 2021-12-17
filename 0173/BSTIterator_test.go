package _0173

import (
    "github.com/raw34/leetcode/runtime"
    "github.com/stretchr/testify/assert"
    "testing"
)

//["BSTIterator", "next", "next", "hasNext", "next", "hasNext", "next", "hasNext", "next", "hasNext"]
//[[[7, 3, 15, null, null, 9, 20]], [], [], [], [], [], [], [], [], []]
//[null, 3, 7, true, 9, true, 15, true, 20, false]
func TestBSTIterator_Case1(t *testing.T) {
    nums := []int{7, 3, 15, 9, 20}
    bst := runtime.NewBinarySearchTree()
    root := bst.BuildFromPreorder(nums)
    itr := Constructor(root)
    assert.Equal(t, 3, itr.Next())
    assert.Equal(t, 7, itr.Next())
    assert.True(t, itr.HasNext())
    assert.Equal(t, 9, itr.Next())
    assert.True(t, itr.HasNext())
    assert.Equal(t, 15, itr.Next())
    assert.True(t, itr.HasNext())
    assert.Equal(t, 20, itr.Next())
    assert.False(t, itr.HasNext())
}
