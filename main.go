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
	graphMap.AddNode("d", graph.Position{
		X: 1.,
		Y: 3.,
	})

	graphMap.AddEdge("a", "b", 0.)
	graphMap.AddEdge("a", "c", 0.)
	graphMap.AddEdge("b", "a", 0.)
	graphMap.AddEdge("c", "a", 0.)

	for id, n := range graphMap.Nodes {
		fmt.Println("node", id, n.Position, n.Edges)
	}

	pathingBFS := graph.BFS{}
	costPath := pathingBFS.CalculateCostPath(&graphMap, "b", "c")

	fmt.Println("BFS: shortest path", costPath)

	graphMap.AddEdge("a", "b", 0.)
	graphMap.AddEdge("a", "c", 2.)
	graphMap.AddEdge("b", "a", 4.)
	graphMap.AddEdge("b", "c", 10.)
	graphMap.AddEdge("c", "a", 0.)
	graphMap.AddEdge("b", "d", 0.)
	graphMap.AddEdge("d", "c", 15.)

	pathingDijkstra := graph.Dijkstra{}
	costPath = pathingDijkstra.CalculateCostPath(&graphMap, "b", "c")

	fmt.Println("Dijkstra: shortest path", costPath)
}
