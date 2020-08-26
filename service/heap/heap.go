// Package heap package supports lfu for impleating heap data structure
package heap

import (
	"container/heap"
	"time"
)

// Data holds key, value and the timestamp when its referenced
type Data struct {
	Key       int
	Value     int
	Timestamp time.Time
}

// PQNode holds the node data structure for priority queue
type PQNode struct {
	NodeData Data
	Freq     int // Frequency(no of times referenced) to maintain the priority
	Index    int // Index in the Priority queue
}

// PriorityQueue is the heap array
type PriorityQueue []*PQNode

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].Freq < pq[j].Freq {
		return true
	}
	// If both nodes have same priority/Frequency select the least recently used
	if pq[i].Freq == pq[j].Freq {
		return pq[i].NodeData.Timestamp.Before(pq[j].NodeData.Timestamp)
	}
	return false
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

// Push inserts node in the heap
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*PQNode)
	item.Index = n
	*pq = append(*pq, item)
}

// Pop removes the most priority element in the heap
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

// Update updates the a particular node and maintains the heap property
func (pq *PriorityQueue) Update(item *PQNode, NodeData Data, Freq int) {
	NodeData.Timestamp = time.Now()
	item.NodeData = NodeData
	item.Freq = Freq
	heap.Fix(pq, item.Index)
}
