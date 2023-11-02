package graph

var CloneGraph = cloneGraph

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	m := map[*Node]*Node{}
	var dfs func(node *Node) *Node
	dfs = func(n *Node) (newN *Node) {
		if n == nil {
			return
		}

		if _, ok := m[n]; ok {
			newN = m[n]
			return
		}

		newN = &Node{
			Val: n.Val,
		}

		m[n] = newN
		neighbors := make([]*Node, len(n.Neighbors))
		for i, ne := range n.Neighbors {
			neighbors[i] = dfs(ne)
		}
		newN.Neighbors = neighbors

		return newN
	}

	return dfs(node)
}
