package _1804

type Trie struct {
    children [26]*Trie
    count    int
    isEnd    bool
}

func Constructor() Trie {
    return Trie{}
}

func (this *Trie) Insert(word string) {
    node := this
    for _, c := range word {
        c -= 'a'
        if node.children[c] == nil {
            node.children[c] = &Trie{}
        }
        node = node.children[c]
        node.count++
    }

    node.isEnd = true
}

func (this *Trie) CountWordsEqualTo(word string) int {
    node := this.countPrefix(word)
    if node != nil && node.isEnd {
        return node.count
    }
    return 0
}

func (this *Trie) countPrefix(prefix string) *Trie {
    node := this
    for _, c := range prefix {
        c -= 'a'
        if node.children[c] == nil {
            return nil
        }
        node = node.children[c]
    }
    return node
}

func (this *Trie) CountWordsStartingWith(prefix string) int {
    node := this.countPrefix(prefix)
    if node != nil {
        return node.count
    }
    return 0
}

func (this *Trie) Erase(word string) {
    node := this
    for _, c := range word {
        c -= 'a'
        if node.children[c] == nil {
            break
        }
        if node.children[c].count == 1 {
            node.children[c] = nil
            break
        }
        node = node.children[c]
        node.count--
    }
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.CountWordsEqualTo(word);
 * param_3 := obj.CountWordsStartingWith(prefix);
 * obj.Erase(word);
 */
