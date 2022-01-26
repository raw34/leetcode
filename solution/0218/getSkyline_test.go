package _0218

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

//输入：buildings = [[2,9,10],[3,7,15],[5,12,12],[15,20,10],[19,24,8]]
//输出：[[2,10],[3,15],[7,12],[12,0],[15,10],[20,8],[24,0]]
func Test_getSkyline_Case1(t *testing.T) {
    buildings := [][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}
    expect := [][]int{{2, 10}, {3, 15}, {7, 12}, {12, 0}, {15, 10}, {20, 8}, {24, 0}}
    assert.Equal(t, expect, getSkyline(buildings))
}

//输入：buildings = [[0,2,3],[2,5,3]]
//输出：[[0,3],[5,0]]
func Test_getSkyline_Case2(t *testing.T) {
    buildings := [][]int{{0, 2, 3}, {2, 5, 3}}
    expect := [][]int{{0, 3}, {5, 0}}
    assert.Equal(t, expect, getSkyline(buildings))
}
