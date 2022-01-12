package _0212

func findWords(board [][]byte, words []string) []string {
    // 初始化字典树
    root := &Trie{}
    for _, word := range words {
        root.insert(word)
    }

    res := make([]string, 0)
    // 标记数组，用于单词排重
    seen := map[string]bool{}
    // 上下左右四个可移动坐标的位移
    dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
    m := len(board)
    n := len(board[0])
    // 回溯方法
    var dfs func(node *Trie, x, y int)
    dfs = func(node *Trie, x, y int) {
        c := board[x][y]
        node = node.children[c-'a']
        // 当前节点不存在时直接退出
        if node == nil {
            return
        }
        // 搜索命中时标记并记录单词
        if node.word != "" && !seen[node.word] {
            seen[node.word] = true
            res = append(res, node.word)
        }
        // 想四个方向回溯
        for _, dir := range dirs {
            nx, ny := x+dir[0], y+dir[1]
            if nx < 0 || ny < 0 || nx >= m || ny >= n || board[nx][ny] == '#' {
                continue
            }
            board[x][y] = '#'
            dfs(node, nx, ny)
            board[x][y] = c
        }
    }
    // 回溯所有可能的单词路径
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            dfs(root, i, j)
        }
    }

    return res
}

type Trie struct {
    children [26]*Trie
    word     string
}

func (this *Trie) insert(word string) {
    node := this
    for _, c := range word {
        c -= 'a'
        if node.children[c] == nil {
            node.children[c] = &Trie{}
        }
        node = node.children[c]
    }
    node.word = word
}
