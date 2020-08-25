package cache

import (
	"container/heap"
	"time"
)

// data holds key, value and the timestamp when its referenced
type data struct {
	key       int
	value     int
	timestamp time.Time
}

// pqNode holds the node data structure for priority queue
type pqNode struct {
	nodeData data
	freq     int // Frequency(no of times referenced) to maintain the priority
	index    int // Index in the Priority queue
}

type priorityQueue []*pqNode

func (pq priorityQueue) Len() int {
	return len(pq)
}
func (pq priorityQueue) Less(i, j int) bool {
	if pq[i].freq < pq[j].freq {
		return true
	}
	// If both nodes have same priority/frequency select the least recently used
	if pq[i].freq == pq[j].freq {
		return pq[i].nodeData.timestamp.Before(pq[j].nodeData.timestamp)
	}
	return false
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*pqNode)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
func (pq *priorityQueue) update(item *pqNode, nodeData data, freq int) {
	nodeData.timestamp = time.Now()
	item.nodeData = nodeData
	item.freq = freq
	heap.Fix(pq, item.index)
}

// cacheLFU hold data structures that implements LFU
type cacheLFU struct {
	keyNodePointerMap map[int]*pqNode
	pq                priorityQueue
}

func (c *cacheLFU) Read(key int) int {
	// If key is present, read the value and update its frequency/priority
	if node, ok := c.keyNodePointerMap[key]; ok {
		return node.nodeData.value
	}
	return -1
}

func (c *cacheLFU) Write(key, value int) {
	// If key is already present, update the value and also update the priority queue
	if node, ok := c.keyNodePointerMap[key]; ok {
		node.nodeData.value = value
		c.pq.update(node, node.nodeData, node.freq+1)
		return
	}
	// If key not present, first check if the length of list is less than capacity of cache else remove the LFU node
	if len(c.pq) >= capacity {
		node := heap.Pop(&c.pq).(*pqNode)
		delete(c.keyNodePointerMap, node.nodeData.key)
	}
	// Pushing the node on heap and inserting the key-Nodepointer in map
	node := &pqNode{data{key, value, time.Now()}, 1, 0}
	heap.Push(&c.pq, node)
	c.keyNodePointerMap[key] = node
}
