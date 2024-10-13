package graph

import (
	"cartographer/priority_queue"
	"container/heap"
)

type Dijkstra struct {
	Queue priority_queue.PriorityQueue
	PathingCommon
}

func (d *Dijkstra) CalculateCostPath(graph *Graph, start string, goal string) (CostPath, error) {
	d.Queue = priority_queue.PriorityQueue{MaxHeap: false}

	item := priority_queue.QueueItem{
		Value:    start,
		Index:    0,
		Priority: 0,
	}

	d.Queue.Push(&item)
	d.Visited = map[string]string{
		start: "",
	}
	d.Costs = map[string]int{
		start: 0,
	}

	for d.Queue.Len() > 0 {
		currentQueueItem := heap.Pop(&d.Queue).(*priority_queue.QueueItem)

		if currentQueueItem.Value == goal {
			break
		}

		current, currentExists := graph.Nodes[currentQueueItem.Value]

		if !currentExists || len(current.Edges) == 0 {
			continue
		}

		for nextNodeId, nextCost := range current.Edges {
			_, nextExists := graph.Nodes[nextNodeId]
			_, hasBeenCosted := d.Costs[nextNodeId]

			newCost := d.Costs[currentQueueItem.Value] + nextCost

			if !nextExists {
				continue
			}

			// Has not yet been costed or the new cost is lower than previous
			if !hasBeenCosted || newCost < d.Costs[nextNodeId] {
				d.Costs[nextNodeId] = newCost

				heap.Push(&d.Queue, &priority_queue.QueueItem{
					Value:    nextNodeId,
					Priority: newCost,
				})

				d.Visited[nextNodeId] = currentQueueItem.Value
			}
		}
	}

	path := d.ReconstructPath(goal)
	// -1 because we don't count the starting node
	moves := len(path) - 1
	cost := d.Costs[goal]

	return CostPath{
		moves,
		cost,
		path,
	}, nil
}
