package graph

import (
	"fmt"
	"slices"
)

type IPathingCommon interface {
	ReconstructPath(goal string)
	CalculateCostPath(graph Graph, start string, goal string, earlyExit bool) CostPath
}

type PathingCommon struct {
	Visited map[string]string
	Costs   map[string]float64
}

func (p *PathingCommon) ReconstructPath(goal string) []string {
	path := []string{}
	current := goal

	for current != "" {
		_, ok := p.Visited[current]
		if !ok {
			break
		}

		path = append(path, current)
		current = p.Visited[current]
	}

	slices.Reverse(path)

	return path
}

type CostPath struct {
	moves int
	cost  float64
	path  []string
}

type Position struct {
	X float64
	Y float64
}

type Node struct {
	Position Position
	Edges    map[string]float64
}

type Graph struct {
	Nodes map[string]Node
}

// Create a node and add it to the graph
func (g *Graph) AddNode(id string, position Position) Node {
	v, exists := g.Nodes[id]

	if exists {
		return v
	}

	edges := make(map[string]float64)

	g.Nodes[id] = Node{
		Position: position,
		Edges:    edges,
	}

	return g.Nodes[id]
}

func (g *Graph) RemoveNode(id string) error {
	_, exists := g.Nodes[id]

	if !exists {
		return fmt.Errorf(`RemoveNode: Node %s does not exist in the graph`, id)
	}

	for _, n := range g.Nodes {
		if len(n.Edges) == 0 {
			continue
		}

		_, ok := n.Edges[id]

		if !ok {
			continue
		}

		delete(n.Edges, id)
	}

	delete(g.Nodes, id)
	return nil
}

func (g *Graph) AddEdge(from string, to string, weight float64) error {
	fromNode, fromExists := g.Nodes[from]
	_, toExists := g.Nodes[to]

	nodes := map[string]bool{
		from: fromExists,
		to:   toExists,
	}

	for id, exists := range nodes {
		if !exists {
			return fmt.Errorf(`AddEdge: Node %s does not exist in the graph`, id)
		}
	}

	fromNode.Edges[to] = weight

	return nil
}
