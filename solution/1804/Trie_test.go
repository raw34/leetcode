package _1804

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

//输入
//["Trie", "insert", "insert", "countWordsEqualTo", "countWordsStartingWith", "erase", "countWordsEqualTo", "countWordsStartingWith", "erase", "countWordsStartingWith"]
//[[], ["apple"], ["apple"], ["apple"], ["app"], ["apple"], ["apple"], ["app"], ["apple"], ["app"]]
//输出
//[null, null, null, 2, 2, null, 1, 1, null, 0]
func TestTrie_Case1(t *testing.T) {
    tree := Constructor()
    tree.Insert("apple")
    tree.Insert("apple")
    assert.Equal(t, 2, tree.CountWordsEqualTo("apple"))
    assert.Equal(t, 2, tree.CountWordsStartingWith("app"))
    tree.Erase("apple")
    assert.Equal(t, 1, tree.CountWordsEqualTo("apple"))
    assert.Equal(t, 1, tree.CountWordsStartingWith("app"))
    tree.Erase("apple")
    assert.Equal(t, 0, tree.CountWordsStartingWith("app"))
}
