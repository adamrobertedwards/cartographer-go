package graph

import (
	"testing"
)

type testCase struct {
	directions    int
	distanceType  DistanceType
	expectedMoves int
	expectedCost  int
}

// TestCalculateCostPathAStar is a unit test for testing each supported distance heuristic for the AStar algorithm.
func TestCalculateCostPathAStar(t *testing.T) {
	testCases := []testCase{
		{directions: 4, distanceType: Manhattan, expectedMoves: 18, expectedCost: 19},
		{directions: 4, distanceType: Chebyshev, expectedMoves: 18, expectedCost: 19},
		{directions: 4, distanceType: Euclidean, expectedMoves: 18, expectedCost: 19},
		{directions: 8, distanceType: Euclidean, expectedMoves: 12, expectedCost: 12},
	}

	for _, test := range testCases {
		grid, err := createGraphGridFromSlice(gridSlice, test.directions)

		if err != nil {
			t.Errorf(`AStar.CalculateCostPath: Graph could not be created: %v`, err.Error())
		}

		pathing := AStar{DistanceType: test.distanceType}
		costPath, costError := pathing.CalculateCostPath(&grid, "0,0", "9,9")

		if costError != nil {
			t.Errorf(`AStar.CalculateCostPath: the cost path could not be calculated with %v : %v`, test.distanceType, costError.Error())
		}

		if costPath.moves != test.expectedMoves {
			t.Errorf(`AStar.CalculateCostPath: the calculated moves were incorrect with %v. expected: %v, got: %v`, test.distanceType, test.expectedMoves, costPath.moves)
		}

		if costPath.cost != test.expectedCost {
			t.Errorf(`AStar.CalculateCostPath: the calculated cost was incorrect with %v. expected: %v, got: %v`, test.distanceType, test.expectedCost, costPath.cost)
		}
	}

}
