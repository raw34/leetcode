package _138

import "github.com/raw34/leetcode/runtime"

var cache map[*runtime.Node]*runtime.Node

func copyRandomList(head *runtime.Node) *runtime.Node {
	cache = map[*runtime.Node]*runtime.Node{}
	return deepCopy(head)
}

func deepCopy(node *runtime.Node) *runtime.Node {
	if node == nil {
		return node
	}

	nodeNew := cache[node]
	if nodeNew != nil {
		return nodeNew
	}

	nodeNew = &runtime.Node{Val: node.Val}
	cache[node] = nodeNew
	nodeNew.Next = deepCopy(node.Next)
	nodeNew.Random = deepCopy(node.Random)

	return nodeNew
}
