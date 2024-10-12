package graph

import (
	"testing"
)

// TestCalculateCostPathAStarManhattan is a unit test for testing Manhattan distance heuristic for the AStar algorithm.
func TestCalculateCostPathAStarManhattan(t *testing.T) {
	grid, err := createGraphGridFromSlice(gridSlice)

	if err != nil {
		t.Errorf(`AStar.CalculateCostPath: Graph could not be created: %v`, err.Error())
	}

	pathing := AStar{DistanceType: Manhattan}
	costPath, costError := pathing.CalculateCostPath(&grid, "0,0", "4,4")
	expectedMoves := 8
	expectedCost := 9.

	if costError != nil {
		t.Errorf(`AStar.CalculateCostPath: the cost path could not be calculated: %v`, costError.Error())
	}

	if costPath.moves != expectedMoves {
		t.Errorf(`AStar.CalculateCostPath: the calculated moves were incorrect. expected: %v, got: %v`, expectedMoves, costPath.moves)
	}

	if costPath.cost != expectedCost {
		t.Errorf(`AStar.CalculateCostPath: the calculated cost was incorrect. expected: %v, got: %v`, expectedCost, costPath.cost)
	}
}
