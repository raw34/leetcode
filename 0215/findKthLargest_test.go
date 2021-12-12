package _0215

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestPriorityQueue_Case1(t *testing.T) {
    arr := []int{3, 2, 1, 5, 6, 4}
    queue := Constructor(arr)
    assert.Equal(t, 6, queue.Pop())
}

//[3,2,1,5,6,4]
//2
func Test_findKthLargest_Case1(t *testing.T) {
    nums := []int{3, 2, 1, 5, 6, 4}
    res := findKthLargest(nums, 2)

    assert.Equal(t, 5, res)
}
