package graph

import (
	"testing"
)

type testCase struct {
	directions    int
	distanceType  DistanceType
	expectedMoves int
	expectedCost  float64
}

// TestCalculateCostPathAStar is a unit test for testing each supported distance heuristic for the AStar algorithm.
func TestCalculateCostPathAStar(t *testing.T) {
	testCases := []testCase{
		{directions: 4, distanceType: Manhattan, expectedMoves: 8, expectedCost: 9.},
		{directions: 8, distanceType: Manhattan, expectedMoves: 6, expectedCost: 6.},
		{directions: 4, distanceType: Euclidean, expectedMoves: 8, expectedCost: 9.},
		{directions: 8, distanceType: Euclidean, expectedMoves: 6, expectedCost: 6.},
	}

	for _, test := range testCases {
		grid, err := createGraphGridFromSlice(gridSlice, test.directions)

		if err != nil {
			t.Errorf(`AStar.CalculateCostPath: Graph could not be created: %v`, err.Error())
		}

		pathing := AStar{DistanceType: test.distanceType}
		costPath, costError := pathing.CalculateCostPath(&grid, "0,0", "4,4")

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
