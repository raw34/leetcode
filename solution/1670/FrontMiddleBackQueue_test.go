package _1670

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

//["FrontMiddleBackQueue", "pushFront", "pushBack", "pushMiddle", "pushMiddle", "popFront", "popMiddle", "popMiddle", "popBack", "popFront"]
//[[], [1], [2], [3], [4], [], [], [], [], []]
//[null, null, null, null, null, 1, 3, 4, 2, -1]
func TestFrontMiddleBackQueue_PopBack(t *testing.T) {
    queue := Constructor()
    queue.PushFront(1)
    queue.PushBack(2)
    queue.PushMiddle(3)
    queue.PushMiddle(4)
    assert.Equal(t, 1, queue.PopFront())
    assert.Equal(t, 3, queue.PopMiddle())
    assert.Equal(t, 4, queue.PopMiddle())
    assert.Equal(t, 2, queue.PopBack())
    assert.Equal(t, -1, queue.PopFront())
}
