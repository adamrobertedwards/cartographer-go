package graph

import (
	"testing"
)

func TestCalculateCostPathBFS(t *testing.T) {
	grid, err := createGraphGridFromSlice(gridSlice)

	if err != nil {
		t.Errorf(`BFS.CalculateCostPath: Graph could not be created: %v`, err.Error())
	}

	pathing := BFS{}
	costPath, costError := pathing.CalculateCostPath(&grid, "0,0", "4,4")
	expectedMoves := 8

	if costError != nil {
		t.Errorf(`BFS.CalculateCostPath: the cost path could not be calculated: %v`, costError.Error())
	}

	if costPath.moves != expectedMoves {
		t.Errorf(`BFS.CalculateCostPath: the calculated moves were incorrect. expected: %v, got: %v`, expectedMoves, costPath.moves)
	}
}
