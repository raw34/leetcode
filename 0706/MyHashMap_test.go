package _706

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

//["MyHashMap", "put", "put", "get", "get", "put", "get", "remove", "get"]
//[[], [1, 1], [2, 2], [1], [3], [2, 1], [2], [2], [2]]
func TestMyHashMap_InOrder1(t *testing.T) {
    hash := Constructor2()
    hash.Put(1, 1)
    hash.Put(2, 2)
    assert.Equal(t, 1, hash.Get(1))
    assert.Equal(t, -1, hash.Get(3))
    hash.Put(2, 1)
    assert.Equal(t, 1, hash.Get(2))
    hash.Remove(2)
    assert.Equal(t, -1, hash.Get(2))
}
