package _0225

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

//["MyStack", "push", "push", "top", "pop", "empty"]
//[[], [1], [2], [], [], []]
//[null, null, null, 2, 2, false]
func TestMyStack_InOrder1(t *testing.T) {
    stack := Constructor()
    stack.Push(1)
    stack.Push(2)
    assert.Equal(t, 2, stack.Top())
    assert.Equal(t, 2, stack.Pop())
    assert.False(t, stack.Empty())
}

//["MyStack", "push", "push", "pop", "top"]
//[[], [1], [2], [], []]
//[null,null,null,2,1]
func TestMyStack_InOrder2(t *testing.T) {
    stack := Constructor()
    stack.Push(1)
    stack.Push(2)
    assert.Equal(t, 2, stack.Pop())
    assert.Equal(t, 1, stack.Top())
}

//["MyStack","push","pop","empty"]
//[[],[1],[],[]]
func TestMyStack_InOrder3(t *testing.T) {
    stack := Constructor()
    stack.Push(1)
    assert.Equal(t, 1, stack.Pop())
    assert.True(t, stack.Empty())
}
