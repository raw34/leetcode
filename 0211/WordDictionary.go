package _0211

type WordDictionary struct {
    children [26]*WordDictionary
    isEnd    bool
}

func Constructor() WordDictionary {
    return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
    node := this
    for _, c := range word {
        c -= 'a'
        if node.children[c] == nil {
            node.children[c] = &WordDictionary{}
        }
        node = node.children[c]
    }
    node.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
    var dfs func(index int, node *WordDictionary) bool
    dfs = func(index int, node *WordDictionary) bool {
        if index == len(word) {
            return node.isEnd
        }

        c := word[index]
        if c != '.' {
            child := node.children[c-'a']
            if child != nil && dfs(index+1, child) {
                return true
            }
        } else {
            for _, child := range node.children {
                if child != nil && dfs(index+1, child) {
                    return true
                }
            }
        }

        return false
    }

    return dfs(0, this)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
