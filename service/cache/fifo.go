package cache

import (
	"github.com/karnadattasai/Cache-Go/service/list"
)

// cacheFIFO hold data structures that implements FIFO
type cacheFIFO struct {
	cacheList         list.List
	keyNodePointerMap map[int]*list.Node
}

func (c *cacheFIFO) Read(key int) int {
	if node, ok := c.keyNodePointerMap[key]; ok {
		return node.P.Value
	}
	return -1
}

func (c *cacheFIFO) Write(key int, value int) {
	// If given key is already present update the value
	if node, ok := c.keyNodePointerMap[key]; ok {
		node.P.Value = value
		return
	}
	// If key not present, first check if the length of list is less than capacity of cache else remove the first entered node
	if c.cacheList.Len() >= capacity {
		delete(c.keyNodePointerMap, c.cacheList.Front().P.Key)
		c.cacheList.Remove(c.cacheList.Front())
	}
	// Add the node to list and insert the key in the map
	c.keyNodePointerMap[key] = c.cacheList.Push(list.Pair{Key: key, Value: value})
}
