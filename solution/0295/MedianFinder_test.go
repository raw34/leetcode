package _0295

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

//["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
//[[], [1], [2], [], [3], []]
//[null, null, null, 1.5, null, 2.0]
func TestMedianFinder_Case1(t *testing.T) {
    finder := Constructor()
    finder.AddNum(1)
    finder.AddNum(2)
    assert.Equal(t, 1.5, finder.FindMedian())
    finder.AddNum(3)
    assert.Equal(t, 2.0, finder.FindMedian())
    finder.AddNum(4)
    assert.Equal(t, 2.5, finder.FindMedian())
}

//["MedianFinder","addNum","findMedian","addNum","findMedian"]
//[[],[2],[],[3],[]]
//[null,null,2.0,null,2.5]
func TestMedianFinder_Case2(t *testing.T) {
    finder := Constructor()
    finder.AddNum(2)
    assert.Equal(t, 2.0, finder.FindMedian())
    finder.AddNum(3)
    assert.Equal(t, 2.5, finder.FindMedian())
}
