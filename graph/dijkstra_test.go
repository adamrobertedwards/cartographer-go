package graph

import (
	"testing"
)

func TestCalculateCostPathDijkstra(t *testing.T) {
	grid, err := createGraphGridFromSlice(gridSlice)

	if err != nil {
		t.Errorf(`Dijkstra.CalculateCostPath: Graph could not be created: %v`, err.Error())
	}

	pathing := Dijkstra{}
	costPath, costError := pathing.CalculateCostPath(&grid, "0,0", "4,4")
	expectedMoves := 8
	expectedCost := 9.

	if costError != nil {
		t.Errorf(`Dijkstra.CalculateCostPath: the cost path could not be calculated: %v`, costError.Error())
	}

	if costPath.moves != expectedMoves {
		t.Errorf(`Dijkstra.CalculateCostPath: the calculated moves were incorrect. expected: %v, got: %v`, expectedMoves, costPath.moves)
	}

	if costPath.cost != expectedCost {
		t.Errorf(`Dijkstra.CalculateCostPath: the calculated cost was incorrect. expected: %v, got: %v`, expectedCost, costPath.cost)
	}
}
