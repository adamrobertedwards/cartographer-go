package graph

import (
	"errors"
	"fmt"
)

var gridSlice [][]int = [][]int{
	{1, 2, 1, 3, 1},
	{1, 3, 1, 1, 1},
	{1, 0, 2, 0, 1},
	{2, 1, 1, 1, 3},
	{1, 1, 0, 1, 1},
}

// createGraphGridFromSlice takes a multidimensional slice, with the values being the cost to visit the node.
func createGraphGridFromSlice(slice [][]int) (Graph, error) {
	if len(slice) == 0 {
		return Graph{}, errors.New("the provided slice was empty")
	}

	nodes := make(map[string]Node)
	grid := Graph{
		Nodes: nodes,
	}

	numRows := len(slice)
	numCols := 0

	for row, cols := range slice {
		numCols = len(cols)

		for col, weight := range cols {
			if weight == 0 {
				continue
			}

			key := fmt.Sprintf("%v,%v", row, col)
			grid.AddNode(key, Position{
				X: float64(row),
				Y: float64(col),
			})
		}
	}

	for key, node := range grid.Nodes {
		col := int(node.Position.Y)
		row := int(node.Position.X)

		if (col - 1) >= 0 {
			prev := slice[row][col-1]

			if prev != 0 {
				grid.AddEdge(key, fmt.Sprintf("%v,%v", row, col-1), float64(prev))
			}
		}

		if (col + 1) <= numCols-1 {
			next := slice[row][col+1]

			if next != 0 {
				grid.AddEdge(key, fmt.Sprintf("%v,%v", row, col+1), float64(next))
			}
		}

		if (row - 1) >= 0 {
			prev := slice[row-1][col]

			if prev != 0 {
				grid.AddEdge(key, fmt.Sprintf("%v,%v", row-1, col), float64(prev))
			}
		}

		if (row + 1) <= numRows-1 {
			next := slice[row+1][col]

			if next != 0 {
				grid.AddEdge(key, fmt.Sprintf("%v,%v", row+1, col), float64(next))
			}
		}
	}

	return grid, nil
}
