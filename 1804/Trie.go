package _1804

type Trie struct {
    children [26]*Trie
    pre      int
    end      int
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
        node.pre++
    }
    node.end++
}

func (this *Trie) CountWordsEqualTo(word string) int {
    node := this
    for _, c := range word {
        c -= 'a'
        if node.children[c] == nil {
            return 0
        }
        node = node.children[c]
    }
    return node.end
}

func (this *Trie) CountWordsStartingWith(prefix string) int {
    node := this
    for _, c := range prefix {
        c -= 'a'
        if node.children[c] == nil {
            return 0
        }
        node = node.children[c]
    }
    return node.pre
}

func (this *Trie) Erase(word string) {
    node := this
    for _, c := range word {
        c -= 'a'
        node = node.children[c]
        node.pre--
    }
    node.end--
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.CountWordsEqualTo(word);
 * param_3 := obj.CountWordsStartingWith(prefix);
 * obj.Erase(word);
 */
