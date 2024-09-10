package main

import (
	"cartographer/graph"
	"fmt"
)

func main() {
	nodes := make(map[string]graph.Node)
	graphMap := graph.Graph{
		Nodes: nodes,
	}

	graphMap.AddNode("a", graph.Position{
		X: 1.,
		Y: 1.,
	})
	graphMap.AddNode("b", graph.Position{
		X: 1.,
		Y: 2.,
	})
	graphMap.AddNode("c", graph.Position{
		X: 2.,
		Y: 2.,
	})

	graphMap.AddEdge("a", "b", 0.)
	graphMap.AddEdge("a", "c", 0.)
	graphMap.AddEdge("b", "a", 0.)
	graphMap.AddEdge("c", "a", 0.)

	for id, n := range graphMap.Nodes {
		fmt.Println("Node", id, n.Position, n.Edges)
	}
}
