package graph

import (
	"fmt"
	"slices"
)

type IPathingCommon interface {
	CalculateCostPath(graph *Graph, start string, goal string) (CostPath, error)
}

// PathingCommon is a struct that contains shared fields; Visited, Costs, which are shared between pathfinding algorithms.
type PathingCommon struct {
	Visited map[string]string
	Costs   map[string]float64
}

// ReconstructPath is a function that returns an array of visited node ids to reach the goal.
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

// Graph is a struct that contains all graph nodes and their edges.
type Graph struct {
	Nodes map[string]Node
}

// AddNode is a function to add a Node to the Graph.
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

// RemoveNode is a function to remove an existing Node from the Graph and any references to the Node in other Node edges.
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

// AddEdge is a function to connect two Nodes together, with an optional weight.
func (g Graph) AddEdge(from string, to string, weight float64) error {
	fromNode, fromExists := g.Nodes[from]
	_, toExists := g.Nodes[to]

	if !fromExists || !toExists {
		return fmt.Errorf(`AddEdge: unable to add edge between nodes because one or both do not exist in the graph`)
	}

	fromNode.Edges[to] = weight

	return nil
}
