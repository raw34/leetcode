package _1028

import (
    "github.com/raw34/leetcode/runtime"
    "github.com/stretchr/testify/assert"
    "testing"
)

func Test_recoverFromPreorder(t *testing.T) {
    str := "1-2--3--4-5--6--7"
    res := recoverFromPreorder(str)
    tree := runtime.BinaryTree{}
    expect := [][]int{{1}, {2, 5}, {3, 4, 6, 7}}

    assert.Equal(t, expect, tree.DisplayLevelOrder1(res))
}
