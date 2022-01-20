package _0211

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

//输入：
//["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
//[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
//输出：
//[null,null,null,null,false,true,true,true]
func TestWordDictionary_AddWord(t *testing.T) {
    tree := Constructor()
    tree.AddWord("bad")
    tree.AddWord("dad")
    tree.AddWord("mad")
    assert.False(t, tree.Search("pad"))
    assert.True(t, tree.Search("bad"))
    assert.True(t, tree.Search(".ad"))
    assert.True(t, tree.Search("b.."))
}
