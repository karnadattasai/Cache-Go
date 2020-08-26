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

// NewLRUCache return new Cache with LRU as replacement policy
func NewLRUCache() Cache {
	c := cacheLRU{}
	c.keyNodePointerMap = make(map[int]*list.Node)
	return &c
}

// NewFIFOCache return new Cache with FIFO as replacement policy
func NewFIFOCache() Cache {
	c := cacheFIFO{}
	c.keyNodePointerMap = make(map[int]*list.Node)
	return &c
}
