package factory

import (
	"errors"

	"github.com/karnadattasai/Cache-Go/service/cache"
)

const (
	// LRU helps in switching cacheType and holds "lru"
	LRU = "lru"
	// FIFO helps in switching cacheType and holds "fifo"
	FIFO = "fifo"
	// LFU helps in switching cacheType and holds "lfu"
	LFU = "lfu"
)

// GetCache returns the Cache interface with requested cache type
func GetCache(cacheType string) (cache.Cache, error) {
	switch cacheType {
	case "lru":
		return cache.NewLRUCache(), nil
	case "fifo":
		return cache.NewFIFOCache(), nil
	case "lfu":
		return cache.NewLFUCache(), nil
	default:
		return nil, errors.New("Undefined cache type. Use one of the following: \"lru\",\"fifo\" or \"lfu\"")
	}
}
