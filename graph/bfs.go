package graph

type BFS struct {
	Queue []string
	PathingCommon
}

func (a *BFS) CalculateCostPath(graph *Graph, start string, goal string) (CostPath, error) {
	a.Queue = []string{start}
	a.Visited = map[string]string{
		start: "",
	}

	for len(a.Queue) > 0 {
		currentNodeId, newQueue := a.Queue[0], a.Queue[1:]
		a.Queue = newQueue

		if currentNodeId == goal {
			break
		}

		current, currentExists := graph.Nodes[currentNodeId]

		if !currentExists || len(current.Edges) == 0 {
			continue
		}

		for nextNodeId := range current.Edges {
			_, nextExists := graph.Nodes[nextNodeId]
			_, alreadyVisited := a.Visited[nextNodeId]

			if !nextExists || alreadyVisited {
				continue
			}

			a.Queue = append(a.Queue, nextNodeId)
			a.Visited[nextNodeId] = currentNodeId
		}
	}

	path := a.ReconstructPath(goal)
	moves := len(path) - 1
	cost := 0

	return CostPath{moves, cost, path}, nil
}
