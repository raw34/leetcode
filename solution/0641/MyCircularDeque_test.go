package _0641

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

//["MyCircularDeque", "insertFront", "getFront", "isEmpty", "deleteFront","insertLast", "getRear", "insertLast", "insertFront", "deleteLast","insertLast", "isEmpty"]
//[[8],[5], [], [],[], [3], [],[7], [7], [], [4], []]
func TestMyCircularDeque_InOrder1(t *testing.T) {
    queue := Constructor(8)
    assert.True(t, queue.InsertFront(5))
    assert.Equal(t, 5, queue.GetFront())
    assert.False(t, queue.IsEmpty())
    assert.True(t, queue.DeleteLast())
    assert.True(t, queue.InsertLast(3))
    assert.Equal(t, 3, queue.GetRear())
    assert.True(t, queue.InsertLast(7))
    assert.True(t, queue.InsertFront(7))
    assert.True(t, queue.DeleteLast())
    assert.True(t, queue.InsertLast(4))
    assert.False(t, queue.IsEmpty())
}

//["MyCircularDeque","insertLast","insertLast","insertFront","insertFront","getRear","isFull","deleteLast","insertFront","getFront"]
//[[3],[1],[2],[3],[4],[],[],[],[4],[]]
func TestMyCircularDeque_InOrder2(t *testing.T) {
    queue := Constructor(3)
    assert.True(t, queue.InsertLast(1))
    assert.True(t, queue.InsertLast(2))
    assert.True(t, queue.InsertFront(3))
    assert.False(t, queue.InsertFront(4))
    assert.Equal(t, 2, queue.GetRear())
    assert.True(t, queue.IsFull())
    assert.True(t, queue.DeleteLast())
    assert.True(t, queue.InsertFront(4))
    assert.Equal(t, 4, queue.GetFront())
}
