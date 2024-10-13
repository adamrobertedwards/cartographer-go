package graph

import (
	"cartographer/priority_queue"
	"container/heap"
	"errors"
)

// AStar is a struct that represents the implementation for the AStar pathfinding algorithm.
type AStar struct {
	Queue        priority_queue.PriorityQueue
	DistanceType DistanceType
	PathingCommon
}

func (a *AStar) CalculateHeuristicCost(from, goal *Node) float64 {
	if a.DistanceType == Manhattan {
		manhattan := &ManhattanDistance{}
		return manhattan.CalculateDistance(from, goal)
	}

	if a.DistanceType == Euclidean {
		euclidean := &EuclideanDistance{}
		return euclidean.CalculateDistance(from, goal)
	}

	return 0.
}

// SetDistanceForHeuristic is a function to set the distance for the heuristic function e.g. manhattan distance.
func (a *AStar) SetDistanceForHeuristic(distance DistanceType) {
	a.DistanceType = distance
}

func (a *AStar) CalculateCostPath(graph *Graph, start string, goal string) (CostPath, error) {
	a.Queue = priority_queue.PriorityQueue{MaxHeap: false}

	item := priority_queue.QueueItem{
		Value:    start,
		Index:    0,
		Priority: 0,
	}

	a.Queue.Push(&item)
	a.Visited = map[string]string{
		start: "",
	}
	a.Costs = map[string]float64{
		start: 0,
	}

	goalNode, goalExists := graph.Nodes[goal]

	if !goalExists {
		return CostPath{0, 0, []string{}}, errors.New("the goal node does not exist")
	}

	for a.Queue.Len() > 0 {
		currentQueueItem := heap.Pop(&a.Queue).(*priority_queue.QueueItem)

		if currentQueueItem.Value == goal {
			break
		}

		current, currentExists := graph.Nodes[currentQueueItem.Value]

		if !currentExists || len(current.Edges) == 0 {
			continue
		}

		for nextNodeId, nextCost := range current.Edges {
			nextNode, nextExists := graph.Nodes[nextNodeId]
			_, hasBeenCosted := a.Costs[nextNodeId]

			newCost := a.Costs[currentQueueItem.Value] + nextCost

			if !nextExists {
				continue
			}

			// Has not yet been costed or the new cost is lower than previous
			if !hasBeenCosted || newCost < a.Costs[nextNodeId] {
				a.Costs[nextNodeId] = newCost

				priority := newCost + a.CalculateHeuristicCost(&nextNode, &goalNode)

				heap.Push(&a.Queue, &priority_queue.QueueItem{
					Value:    nextNodeId,
					Priority: priority,
				})

				a.Visited[nextNodeId] = currentQueueItem.Value
			}
		}
	}

	path := a.ReconstructPath(goal)
	moves := len(path) - 1
	cost := a.Costs[goal]

	return CostPath{
		moves,
		cost,
		path,
	}, nil
}
