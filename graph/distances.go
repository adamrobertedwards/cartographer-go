package graph

import (
	"math"
)

type DistanceType string

const (
	Manhattan DistanceType = "manhattan"
	Euclidean DistanceType = "euclidean"
)

type DistanceCommon interface {
	CalculateDistance(from, to *Node) float64
}

type ManhattanDistance struct{}

func (e ManhattanDistance) CalculateDistance(from, goal *Node) float64 {
	dx := math.Abs(from.Position.X - goal.Position.X)
	dy := math.Abs(from.Position.Y - goal.Position.Y)

	return dx + dy
}

type EuclideanDistance struct{}

func (e EuclideanDistance) CalculateDistance(from, goal *Node) float64 {
	dx := math.Abs(from.Position.X - goal.Position.X)
	dy := math.Abs(from.Position.Y - goal.Position.Y)

	return math.Sqrt((dx * dx) + (dy * dy))
}
