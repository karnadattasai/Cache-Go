package factory

import (
	"errors"

	"github.com/karnadattasai/Cache-Go/service/cache"
)

// GetCache returns the Cache interface with requested cache type
func GetCache(cacheType string) (cache.Cache, error) {
	switch cacheType {
	case "lru":
		return cache.NewLRUCache(), nil
	case "fifo":
		return cache.NewFIFOCache(), nil
	default:
		return nil, errors.New("Undefined cache type. Use one of the following: \"lru\",\"fifo\" or \"lfu\"")
	}
}
