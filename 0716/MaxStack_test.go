package _716

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

//["MaxStack", "push", "push", "push", "top", "popMax", "top", "peekMax", "pop", "top"]
//[[], [5], [1], [5], [], [], [], [], [], []]
//[null, null, null, null, 5, 5, 1, 5, 1, 5]
func TestMaxStack_Case1(t *testing.T) {
    stack := Constructor()
    stack.Push(5)
    stack.Push(1)
    stack.Push(5)
    assert.Equal(t, 5, stack.Top())
    assert.Equal(t, 5, stack.PopMax())
    assert.Equal(t, 1, stack.Top())
    assert.Equal(t, 5, stack.PeekMax())
    assert.Equal(t, 1, stack.Pop())
    assert.Equal(t, 5, stack.Top())
}

//["MaxStack","push","push","popMax","peekMax"]
//[[],[5],[1],[],[]]
func TestMaxStack_Case2(t *testing.T) {
    stack := Constructor()
    stack.Push(5)
    stack.Push(1)
    assert.Equal(t, 5, stack.PopMax())
    assert.Equal(t, 1, stack.PeekMax())
}
