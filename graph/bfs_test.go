package graph

import (
	"testing"
)

type testCaseBFS struct {
	directions    int
	expectedMoves int
}

func TestCalculateCostPathBFS(t *testing.T) {
	testCases := []testCaseBFS{
		{directions: 4, expectedMoves: 8},
		{directions: 8, expectedMoves: 4},
	}

	for _, test := range testCases {
		grid, err := createGraphGridFromSlice(gridSlice, test.directions)

		if err != nil {
			t.Errorf(`BFS.CalculateCostPath: Graph could not be created: %v`, err.Error())
		}

		pathing := BFS{}
		costPath, costError := pathing.CalculateCostPath(&grid, "0,0", "4,4")

		if costError != nil {
			t.Errorf(`BFS.CalculateCostPath: the cost path could not be calculated: %v`, costError.Error())
		}

		if costPath.moves != test.expectedMoves {
			t.Errorf(`BFS.CalculateCostPath: the calculated moves were incorrect. expected: %v, got: %v`, test.expectedMoves, costPath.moves)
		}
	}
}
