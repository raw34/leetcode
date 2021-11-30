package _450

import (
	"github.com/raw34/leetcode/runtime"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_deleteNode_OneNode(t *testing.T) {
	root := &runtime.TreeNode{Val: 0}

	res := deleteNode(root, 0)

	assert.Nil(t, nil, res)
}