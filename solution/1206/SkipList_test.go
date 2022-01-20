package _1206

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

//["Skiplist", "add", "add", "add", "search","add", "search", "erase", "erase", "search"]
//[[], [1], [2], [3], [0], [4], [1], [0], [1], [1]]
//[null, null, null, null, false,null, true, false, true, false]
func TestSkiplist_Case1(t *testing.T) {
    list := Constructor()
    list.Add(1)
    list.Add(2)
    list.Add(3)
    assert.False(t, list.Search(0))
    list.Add(4)
    assert.True(t, list.Search(1))
    assert.False(t, list.Erase(0))
    assert.True(t, list.Erase(1))
    assert.False(t, list.Search(1))
}

//["Skiplist", "add", "add", "add", "add","search", "erase", "search", "search", "search"]
//[[], [0], [5], [2], [1], [0], [5], [2], [3], [2]]
//[null,null,null,null,null,true,true,true,false,true]
func TestSkiplist_Case2(t *testing.T) {
    list := Constructor()
    list.Add(1)
    list.Add(0)
    list.Add(5)
    list.Add(2)
    list.Add(1)
    assert.True(t, list.Search(0))
    assert.True(t, list.Erase(5))
    assert.True(t, list.Search(2))
    assert.False(t, list.Search(3))
    assert.True(t, list.Search(2))
}
