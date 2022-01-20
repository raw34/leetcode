package _0622

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

//["MyCircularQueue", "enQueue", "enQueue", "enQueue", "enQueue","Rear", "isFull", "deQueue", "enQueue", "Rear"]
//[[3], [1],[2], [3], [4], [], [], [], [4], []]
func TestMyCircularQueue_InOrder1(t *testing.T) {
    queue := Constructor(3)
    assert.True(t, queue.EnQueue(1))
    assert.True(t, queue.EnQueue(2))
    assert.True(t, queue.EnQueue(3))
    assert.False(t, queue.EnQueue(4))
    assert.Equal(t, 3, queue.Rear())
    assert.True(t, queue.IsFull())
    assert.True(t, queue.DeQueue())
    assert.True(t, queue.EnQueue(4))
    assert.Equal(t, 4, queue.Rear())
}

//["MyCircularQueue","enQueue","Rear","Rear","deQueue","enQueue","Rear","deQueue","Front","deQueue","deQueue","deQueue"]
//[[6],[6],[],[],[],[5],[],[],[],[],[],[]]
func TestMyCircularQueue_InOrder2(t *testing.T) {
    queue := Constructor(6)
    assert.True(t, queue.EnQueue(6))
    assert.Equal(t, 6, queue.Rear())
    assert.Equal(t, 6, queue.Rear())
    assert.True(t, queue.DeQueue())
    assert.True(t, queue.EnQueue(5))
    assert.Equal(t, 5, queue.Rear())
    assert.True(t, queue.DeQueue())
    assert.Equal(t, -1, queue.Front())
    assert.False(t, queue.DeQueue())
    assert.False(t, queue.DeQueue())
    assert.False(t, queue.DeQueue())
}

//["MyCircularQueue","enQueue","enQueue","enQueue","enQueue","deQueue","deQueue","isEmpty","isEmpty","Rear","Rear","deQueue"]
//[[8],[3],[9],[5],[0],[],[],[],[],[],[],[]]
func TestMyCircularQueue_InOrder3(t *testing.T) {
    queue := Constructor(8)
    assert.True(t, queue.EnQueue(3))
    assert.True(t, queue.EnQueue(9))
    assert.True(t, queue.EnQueue(5))
    assert.True(t, queue.EnQueue(0))
    assert.True(t, queue.DeQueue())
    assert.True(t, queue.DeQueue())
    assert.False(t, queue.IsEmpty())
    assert.False(t, queue.IsEmpty())
    assert.Equal(t, 0, queue.Rear())
    assert.Equal(t, 0, queue.Rear())
    assert.True(t, queue.DeQueue())
}

//["MyCircularQueue","enQueue","enQueue","Front","enQueue","deQueue","enQueue","enQueue","Rear","isEmpty","Front","deQueue"]
//[[2],[8],[8],[],[4],[],[1],[1],[],[],[],[]]
func TestMyCircularQueue_InOrder4(t *testing.T) {
    queue := Constructor(2)
    assert.True(t, queue.EnQueue(8))
    assert.True(t, queue.EnQueue(8))
    assert.Equal(t, 8, queue.Front())
    assert.False(t, queue.EnQueue(4))
    assert.True(t, queue.DeQueue())
    assert.True(t, queue.EnQueue(1))
    assert.False(t, queue.EnQueue(1))
    assert.Equal(t, 1, queue.Rear())
    assert.False(t, queue.IsEmpty())
    assert.Equal(t, 8, queue.Front())
    assert.True(t, queue.DeQueue())
}
