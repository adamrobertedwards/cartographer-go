package graph

import (
	"math"
)

type DistanceType string

const (
	Manhattan DistanceType = "manhattan"
	Chebyshev DistanceType = "chebyshev"
	Euclidean DistanceType = "euclidean"
)

type DistanceCommon interface {
	CalculateDistance(from, to *Node) float64
}

type ManhattanDistance struct{}

// CalculateDistance on the ManhattanDistance struct is a function that finds the shortest distance between two points using 4-way movement.
func (e *ManhattanDistance) CalculateDistance(from, goal *Node) int {
	dx := math.Abs(from.Position.X - goal.Position.X)
	dy := math.Abs(from.Position.Y - goal.Position.Y)

	return int(dx + dy)
}

type ChebyshevDistance struct{}

// CalculateDistance on the ChebyshevDistance struct is a function that finds the shortest distance between two points using 8-way movement.
func (e *ChebyshevDistance) CalculateDistance(from, goal *Node) int {
	dx := math.Abs(from.Position.X - goal.Position.X)
	dy := math.Abs(from.Position.Y - goal.Position.Y)

	return int(math.Max(dx, dy))
}

type EuclideanDistance struct{}

// CalculateDistance on the EuclideanDistance struct is a function that finds the shortest line between two points using euclidean distance e.g. straight line distance.
func (e *EuclideanDistance) CalculateDistance(from, goal *Node) int {
	dx := math.Abs(from.Position.X - goal.Position.X)
	dy := math.Abs(from.Position.Y - goal.Position.Y)

	return int(math.Sqrt((dx * dx) + (dy * dy)))
}
