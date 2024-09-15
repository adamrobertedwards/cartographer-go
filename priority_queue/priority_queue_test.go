package priority_queue

import (
	"container/heap"
	"testing"
)

func setupQueueForTests(maxHeap bool) PriorityQueue {
	items := []*QueueItem{
		{
			Value:    "apple",
			Priority: 3,
			Index:    0,
		},
		{
			Value:    "banana",
			Priority: 2,
			Index:    1,
		},
		{
			Value:    "orange",
			Priority: 1,
			Index:    2,
		},
	}
	a := PriorityQueue{
		items,
		maxHeap,
	}

	return a
}

func TestPriorityQueuePop(t *testing.T) {
	tests := []struct {
		name     string
		MaxHeap  bool
		expected *QueueItem
	}{
		{
			name:    "Priority queue with min-heap",
			MaxHeap: false,
			expected: &QueueItem{
				Value:    "orange",
				Priority: 1,
				Index:    2,
			},
		},
		{
			name:    "Priority queue with max-heap",
			MaxHeap: true,
			expected: &QueueItem{
				Value:    "apple",
				Priority: 3,
				Index:    0,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			queue := setupQueueForTests(test.MaxHeap)
			heap.Init(&queue)

			popped := heap.Pop(&queue).(*QueueItem)

			if popped.Value != test.expected.Value || popped.Priority != test.expected.Priority {
				t.Errorf(`PriorityQueue.Pop(): should return the highest or lowest priority item = %v but got = %v when maxHeap = %v`, test.expected, popped, test.MaxHeap)
			}

			if queue.Len() != 2 {
				t.Errorf(`PriorityQueue.Pop(): should reduce the queue length by one item but got: %v`, queue.Len())
			}

		})
	}
}

func TestPriorityQueuePush(t *testing.T) {
	queue := setupQueueForTests(false)

	heap.Init(&queue)

	item := &QueueItem{
		Value:    "kiwi",
		Priority: 2,
		Index:    0,
	}

	heap.Push(&queue, item)

	queueLength := queue.Len()

	if queueLength != 4 {
		t.Errorf(`PriorityQueue.Push(): should increase the queue length by one item but got: %v`, queue.Len())
	}

	if queue.items[queueLength-1] != item {
		t.Errorf(`PriorityQueue.Push(): should have added the new item to the queue but got: %v`, queue.items[queueLength-1])
	}
}
