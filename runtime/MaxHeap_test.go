package runtime

import (
    "container/heap"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestMaxHeap_Case1(t *testing.T) {
    arr := []int{3, 2, 1, 5, 6, 4}
    queue := NewMaxHeap(arr)
    assert.Equal(t, 6, heap.Pop(queue))
    assert.Equal(t, 5, heap.Pop(queue))
    assert.Equal(t, 4, heap.Pop(queue))
    assert.Equal(t, 3, heap.Pop(queue))
    assert.Equal(t, 2, heap.Pop(queue))
    assert.Equal(t, 1, heap.Pop(queue))
}
