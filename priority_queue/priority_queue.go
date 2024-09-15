package priority_queue

type QueueItem struct {
	// It's more efficient to store the index as a field vs traversing the heap to get the index explicitly
	Index    int
	Priority float64
	Value    string
}

// Array of pointers to queue items
type PriorityQueue struct {
	items   []*QueueItem
	MaxHeap bool
}

// Implement the functions of heap

func (pq PriorityQueue) Len() int {
	return len(pq.items)
}

func (pq PriorityQueue) Less(a, b int) bool {
	if pq.MaxHeap {
		return pq.items[a].Priority > pq.items[b].Priority
	}
	return pq.items[a].Priority < pq.items[b].Priority
}

func (pq PriorityQueue) Swap(a, b int) {
	pq.items[a], pq.items[b] = pq.items[b], pq.items[a]
	pq.items[a].Index = b
	pq.items[a].Index = a
}

// Take a pointer to PriorityQueue, so we aren't working with a copy
func (pq *PriorityQueue) Push(i interface{}) {
	item := i.(*QueueItem)
	item.Index = pq.Len()
	pq.items = append(pq.items, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	n := pq.Len()
	item := pq.items[n-1]
	pq.items[n-1] = nil
	pq.items = pq.items[:n-1]

	return item
}
