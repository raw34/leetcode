package _0208

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

//输入
//["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
//[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
//输出
//[null, null, true, false, true, null, true]
func TestTrie_Case1(t *testing.T) {
    tree := Constructor()
    tree.Insert("apple")
    assert.True(t, tree.Search("apple"))
    assert.False(t, tree.Search("app"))
    assert.True(t, tree.StartsWith("app"))
    tree.Insert("app")
    assert.True(t, tree.Search("app"))
}
