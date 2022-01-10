package _0138

type Node struct {
    Val    int
    Next   *Node
    Random *Node
}

var cache map[*Node]*Node

func copyRandomList(head *Node) *Node {
    cache = map[*Node]*Node{}
    return deepCopy(head)
}

func deepCopy(node *Node) *Node {
    if node == nil {
        return node
    }

    nodeNew := cache[node]
    if nodeNew != nil {
        return nodeNew
    }

    nodeNew = &Node{Val: node.Val}
    cache[node] = nodeNew
    nodeNew.Next = deepCopy(node.Next)
    nodeNew.Random = deepCopy(node.Random)

    return nodeNew
}
