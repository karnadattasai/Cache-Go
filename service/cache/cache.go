//Package cache Provides LRU, LFU, FIFU Cache
package cache

import "github.com/karnadattasai/Cache-Go/service/list"

//capacity is the maximum size of the Cache
const capacity = 3

// Cache interface contains methods read, write and new methods
type Cache interface {
	Read(key int) (value int) // reads and returns a value at given key
	Write(key int, value int) // writes the given value at given key
}

// cacheAbstraction abstracts lru and fifo cache
type cacheAbstraction struct {
	cacheList         list.List
	keyNodePointerMap map[int]*list.Node
}

func (c *cacheAbstraction) Read(key int) int {
	if node, ok := c.keyNodePointerMap[key]; ok {
		return node.P.Value
	}
	return -1
}

// NewLRUCache return new Cache with LRU as replacement policy
func NewLRUCache() Cache {
	c := cacheLRU{cacheAbstraction{}}
	c.keyNodePointerMap = make(map[int]*list.Node)
	return &c
}

// NewFIFOCache return new Cache with FIFO as replacement policy
func NewFIFOCache() Cache {
	c := cacheFIFO{cacheAbstraction{}}
	c.keyNodePointerMap = make(map[int]*list.Node)
	return &c
}
