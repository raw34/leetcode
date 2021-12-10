package _0641

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

//["MyCircularDeque", "insertFront", "getFront", "isEmpty", "deleteFront","insertLast", "getRear", "insertLast", "insertFront", "deleteLast","insertLast", "isEmpty"]
//[[8],[5], [], [],[], [3], [],[7], [7], [], [4], []]
//[null,true,5,false,true,true,3,true,true,true,true,false]
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
