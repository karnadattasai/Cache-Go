//Package cache Provides LRU, LFU, FIFU Cache
package cache

import (
	"container/heap"
)

//capacity is the maximum size of the Cache
const capacity = 3

// Cache interface contains methods read, write and new methods
type Cache interface {
	Read(key int) (value int) // reads and returns a value at given key
	Write(key int, value int) // writes the given value at given key
}

// NewLFUCache return new Cache with LFU as replacement policy
func NewLFUCache() Cache {
	c := cacheLFU{}
	c.pq = make(priorityQueue, 0)
	heap.Init(&c.pq)
	c.keyNodePointerMap = make(map[int]*pqNode)
	return &c
}
