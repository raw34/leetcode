package _0705

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

//["MyHashSet2","add","add","contains","contains","add","contains","remove","contains"]
//[[],[1],[2],[1],[3],[2],[2],[2],[2]]
func TestMyHashSet_InOrder1(t *testing.T) {
    set := Constructor2()
    set.Add(1)
    set.Add(2)
    assert.True(t, set.Contains(1))
    assert.False(t, set.Contains(3))
    set.Add(2)
    assert.True(t, set.Contains(2))
    set.Remove(2)
    assert.False(t, set.Contains(2))
}
