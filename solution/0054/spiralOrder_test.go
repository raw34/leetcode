package _0054

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,3,6,9,8,7,4,5]
func Test_spiralOrder_Case1(t *testing.T) {
    matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
    assert.Equal(t, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}, spiralOrder(matrix))
}

//输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//输出：[1,2,3,4,8,12,11,10,9,5,6,7]
func Test_spiralOrder_Case2(t *testing.T) {
    matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
    assert.Equal(t, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}, spiralOrder(matrix))
}

//[[3],[2]]
//[3,2]
func Test_spiralOrder_Case3(t *testing.T) {
    matrix := [][]int{{3}, {2}}
    assert.Equal(t, []int{3, 2}, spiralOrder(matrix))
}

//[[1]]
//[1]
func Test_spiralOrder_Case4(t *testing.T) {
    matrix := [][]int{{1}}
    assert.Equal(t, []int{1}, spiralOrder(matrix))
}

//[[2,5],[8,4],[0,-1]]
//[2,5,4,-1,0,8]
func Test_spiralOrder_Case5(t *testing.T) {
    matrix := [][]int{{2, 5}, {8, 4}, {0, -1}}
    assert.Equal(t, []int{2, 5, 4, -1, 0, 8}, spiralOrder(matrix))
}
