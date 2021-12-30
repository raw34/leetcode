package _1473

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

//[0,0,0,0,0]
//[[1,10],[10,1],[10,1],[1,10],[5,1]]
//5
//2
//3
func Test_minCost_Case1(t *testing.T) {
    houses := []int{0, 0, 0, 0, 0}
    cost := [][]int{{1, 10}, {10, 1}, {10, 1}, {1, 10}, {5, 1}}
    m := 5
    n := 2
    target := 3
    res := minCost(houses, cost, m, n, target)
    assert.Equal(t, 9, res)
}

//[0,2,1,2,0]
//[[1,10],[10,1],[10,1],[1,10],[5,1]]
//5
//2
//3
func Test_minCost_Case2(t *testing.T) {
    houses := []int{0, 2, 1, 2, 0}
    cost := [][]int{{1, 10}, {10, 1}, {10, 1}, {1, 10}, {5, 1}}
    m := 5
    n := 2
    target := 3
    res := minCost(houses, cost, m, n, target)
    assert.Equal(t, 11, res)
}

//[2,2,1]
//[[1,1],[3,4],[4,2]]
//3
//2
//2
func Test_minCost_Case3(t *testing.T) {
    houses := []int{2, 2, 1}
    cost := [][]int{{1, 1}, {3, 4}, {4, 2}}
    m := 3
    n := 2
    target := 2
    res := minCost(houses, cost, m, n, target)
    assert.Equal(t, 0, res)
}

//[3,1,2,3]
//[[1,1,1],[1,1,1],[1,1,1],[1,1,1]]
//4
//3
//3
func Test_minCost_Case4(t *testing.T) {
    houses := []int{3, 1, 2, 3}
    cost := [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}}
    m := 4
    n := 3
    target := 3
    res := minCost(houses, cost, m, n, target)
    assert.Equal(t, -1, res)
}

//[0,0,0,0,0]
//[[1,10],[10,1],[1,10],[10,1],[1,10]]
//5
//2
//5
func Test_minCost_Case5(t *testing.T) {
    houses := []int{0, 0, 0, 0, 0}
    cost := [][]int{{1, 10}, {10, 1}, {1, 10}, {10, 1}, {1, 10}}
    m := 5
    n := 2
    target := 5
    res := minCost(houses, cost, m, n, target)
    assert.Equal(t, 5, res)
}
