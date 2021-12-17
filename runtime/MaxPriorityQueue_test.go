package runtime

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestPriorityQueue_Case1(t *testing.T) {
    arr := []int{3, 2, 1, 5, 6, 4}
    queue := NewMaxPriorityQueue(arr)
    assert.Equal(t, 6, queue.Pop())
    assert.Equal(t, 5, queue.Pop())
    assert.Equal(t, 4, queue.Pop())
    assert.Equal(t, 3, queue.Pop())
    assert.Equal(t, 2, queue.Pop())
    assert.Equal(t, 1, queue.Pop())
}
