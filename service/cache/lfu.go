package cache

import (
	"container/heap"
	"time"

	heapForLFU "github.com/karnadattasai/Cache-Go/service/heap"
)

// cacheLFU hold data structures that implements LFU
type cacheLFU struct {
	keyNodePointerMap map[int]*heapForLFU.PQNode
	pq                heapForLFU.PriorityQueue
}

func (c *cacheLFU) Read(key int) int {
	// If key is present, read the value and update its frequency/priority
	if node, ok := c.keyNodePointerMap[key]; ok {
		return node.NodeData.Value
	}
	return -1
}

func (c *cacheLFU) Write(key, value int) {
	// If key is already present, update the Value and also update the priority queue
	if node, ok := c.keyNodePointerMap[key]; ok {
		node.NodeData.Value = value
		c.pq.Update(node, node.NodeData, node.Freq+1)
		return
	}
	// If key not present, first check if the length of list is less than capacity of cache else remove the LFU node
	if len(c.pq) >= capacity {
		node := heap.Pop(&c.pq).(*heapForLFU.PQNode)
		delete(c.keyNodePointerMap, node.NodeData.Key)
	}
	// Pushing the node on heap and inserting the key-Nodepointer in map
	node := &heapForLFU.PQNode{NodeData: heapForLFU.Data{Key: key, Value: value, Timestamp: time.Now()}, Freq: 1, Index: 0}
	heap.Push(&c.pq, node)
	c.keyNodePointerMap[key] = node
}
