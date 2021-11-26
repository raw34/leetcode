package _133

type Node struct {
	Val int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}

	cache := map[*Node]*Node{}
	queue := []*Node{node}
	cache[node] = &Node{
		Val: node.Val,
		Neighbors: []*Node{},
	}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, neighbor := range curr.Neighbors {
			if cache[neighbor] == nil {
				cache[neighbor] = &Node{
					Val: neighbor.Val,
					Neighbors: []*Node{},
				}
				queue = append(queue, neighbor)
			}
			cache[curr].Neighbors = append(cache[curr].Neighbors, cache[neighbor])
		}
	}

	return cache[node]
}