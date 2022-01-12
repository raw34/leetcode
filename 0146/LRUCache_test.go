package _0146

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

//输入
//["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
//输出
//[null, null, null, 1, null, -1, null, -1, 3, 4]
func TestLRUCache_Case1(t *testing.T) {
    cache := Constructor(2)
    cache.Put(1, 1)
    cache.Put(2, 2)
    assert.Equal(t, 1, cache.Get(1))
    cache.Put(3, 3)
    assert.Equal(t, -1, cache.Get(2))
    cache.Put(4, 4)
    assert.Equal(t, -1, cache.Get(1))
    assert.Equal(t, 3, cache.Get(3))
    assert.Equal(t, 4, cache.Get(4))
}

//输入：
//["LRUCache","put","put","get","put","get","put","get","get","get"]
//[[2],[1,0],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
//输出：
//[null,null,null,0,null,-1,null,-1,3,4]
func TestLRUCache_Case2(t *testing.T) {
    cache := Constructor(2)
    cache.Put(1, 0)
    cache.Put(2, 2)
    assert.Equal(t, 0, cache.Get(1))
    cache.Put(3, 3)
    assert.Equal(t, -1, cache.Get(2))
    cache.Put(4, 4)
    assert.Equal(t, -1, cache.Get(1))
    assert.Equal(t, 3, cache.Get(3))
    assert.Equal(t, 4, cache.Get(4))
}
