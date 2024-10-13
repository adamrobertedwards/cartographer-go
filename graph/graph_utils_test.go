package graph

import (
	"errors"
	"fmt"
	"math"
)

var gridSlice [][]int = [][]int{
	{1, 2, 1, 3, 1, 1, 2, 0, 1, 1},
	{1, 2, 1, 1, 1, 2, 1, 1, 1, 0},
	{1, 0, 2, 0, 1, 1, 1, 2, 3, 1},
	{2, 1, 1, 1, 3, 2, 1, 1, 1, 1},
	{1, 1, 0, 1, 1, 0, 2, 1, 1, 1},
	{1, 2, 1, 3, 1, 1, 2, 1, 1, 1},
	{1, 2, 1, 1, 1, 2, 1, 1, 1, 0},
	{1, 0, 2, 1, 1, 1, 1, 2, 1, 1},
	{2, 1, 1, 1, 3, 2, 1, 1, 1, 1},
	{1, 1, 0, 1, 1, 0, 2, 1, 1, 1},
}

// createGraphGridFromSlice takes a multidimensional slice, with the values being the cost to visit the node.
func createGraphGridFromSlice(slice [][]int, directions int) (Graph, error) {
	if len(slice) == 0 {
		return Graph{}, errors.New("the provided slice was empty")
	}

	// On a grid, only support 4 or 8 way movement
	if directions != 4 && directions != 8 {
		return Graph{}, errors.New("the provided value for directions must be 4 or 8")
	}

	nodes := make(map[string]Node)
	grid := Graph{
		Nodes: nodes,
	}

	numRows := len(slice)
	numCols := 0
	diagonalCost := int(math.Sqrt(2))

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
				grid.AddEdge(key, fmt.Sprintf("%v,%v", row, col-1), prev)
			}
		}

		// Add edges for 4 way movement
		if (col + 1) <= numCols-1 {
			next := slice[row][col+1]

			if next != 0 {
				grid.AddEdge(key, fmt.Sprintf("%v,%v", row, col+1), next)
			}
		}

		if (row - 1) >= 0 {
			prev := slice[row-1][col]

			if prev != 0 {
				grid.AddEdge(key, fmt.Sprintf("%v,%v", row-1, col), prev)
			}
		}

		if (row + 1) <= numRows-1 {
			next := slice[row+1][col]

			if next != 0 {
				grid.AddEdge(key, fmt.Sprintf("%v,%v", row+1, col), next)
			}
		}

		// Handle edges for 8 way movement
		if directions == 8 {
			if (col-1) >= 0 && (row-1) >= 0 {
				node := slice[row-1][col-1]

				if node != 0 {
					grid.AddEdge(key, fmt.Sprintf("%v,%v", row-1, col-1), node*diagonalCost)
				}
			}

			if (col+1) <= numCols-1 && (row-1) >= 0 {
				node := slice[row-1][col+1]

				if node != 0 {
					grid.AddEdge(key, fmt.Sprintf("%v,%v", row-1, col+1), node*diagonalCost)
				}
			}

			if (col-1) >= 0 && (row+1) <= numRows-1 {
				node := slice[row+1][col-1]

				if node != 0 {
					grid.AddEdge(key, fmt.Sprintf("%v,%v", row+1, col-1), node*diagonalCost)
				}
			}

			if (col+1) <= numCols-1 && (row+1) <= numRows-1 {
				node := slice[row+1][col+1]

				if node != 0 {
					grid.AddEdge(key, fmt.Sprintf("%v,%v", row+1, col+1), node*diagonalCost)
				}
			}
		}
	}

	return grid, nil
}
