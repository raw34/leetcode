package _0215

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

//[3,2,1,5,6,4]
//2
//5
func Test_findKthLargest_Case1(t *testing.T) {
    nums := []int{3, 2, 1, 5, 6, 4}
    res := findKthLargest(nums, 2)

    assert.Equal(t, 5, res)
}
