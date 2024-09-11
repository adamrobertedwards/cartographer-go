package graph

import (
	"slices"
)

type BFS struct {
	Pathing
}

func (a *BFS) ReconstructPath(goal string) []string {
	path := []string{}
	current := goal

	for current != "" {
		_, ok := a.Visited[current]
		if !ok {
			break
		}

		path = append(path, current)
		current = a.Visited[current]
	}

	slices.Reverse(path)

	return path
}

func (a *BFS) CalculateCostPath(graph *Graph, start string, goal string) CostPath {
	a.Queue = []string{start}
	a.Visited = map[string]string{
		start: "",
	}
	moves := 0

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

			moves++
		}
	}

	path := a.ReconstructPath(goal)
	cost := 0.

	return CostPath{moves, cost, path}
}
