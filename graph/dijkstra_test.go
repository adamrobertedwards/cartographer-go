package graph

import (
	"testing"
)

type testCaseDijkstra struct {
	directions    int
	expectedMoves int
	expectedCost  float64
}

func TestCalculateCostPathDijkstra(t *testing.T) {
	testCases := []testCaseDijkstra{
		{directions: 4, expectedMoves: 8, expectedCost: 9.},
		{directions: 8, expectedMoves: 6, expectedCost: 6.},
	}

	for _, test := range testCases {
		grid, err := createGraphGridFromSlice(gridSlice, test.directions)

		if err != nil {
			t.Errorf(`Dijkstra.CalculateCostPath: Graph could not be created: %v`, err.Error())
		}

		pathing := Dijkstra{}
		costPath, costError := pathing.CalculateCostPath(&grid, "0,0", "4,4")

		if costError != nil {
			t.Errorf(`Dijkstra.CalculateCostPath: the cost path could not be calculated: %v`, costError.Error())
		}

		if costPath.moves != test.expectedMoves {
			t.Errorf(`Dijkstra.CalculateCostPath: the calculated moves were incorrect. expected: %v, got: %v`, test.expectedMoves, costPath.moves)
		}

		if costPath.cost != test.expectedCost {
			t.Errorf(`Dijkstra.CalculateCostPath: the calculated cost was incorrect. expected: %v, got: %v`, test.expectedCost, costPath.cost)
		}
	}
}
