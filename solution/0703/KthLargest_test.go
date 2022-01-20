package _0703

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

//["KthLargest", "add", "add", "add", "add", "add"]
//[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
//[null, 4, 5, 5, 8, 8]
func TestKthLargest_Case1(t *testing.T) {
    k := 3
    nums := []int{4, 5, 8, 2}
    kth := Constructor(k, nums)
    assert.Equal(t, 4, kth.Add(3))
    assert.Equal(t, 5, kth.Add(5))
    assert.Equal(t, 5, kth.Add(10))
    assert.Equal(t, 8, kth.Add(9))
    assert.Equal(t, 8, kth.Add(4))
}
