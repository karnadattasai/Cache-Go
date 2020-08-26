package factory

import (
	"errors"
	"testing"

	"github.com/karnadattasai/Cache-Go/service/cache"

	"github.com/stretchr/testify/assert"
)

func Test_GetCache(t *testing.T) {

	tests := []struct {
		name      string
		cacheType string
		want      cache.Cache
		err       error
	}{
		{
			name:      "[Success] lru check",
			cacheType: "lru",
			want:      cache.NewLRUCache(),
			err:       nil,
		},
		{
			name:      "[Success] fifo cache",
			cacheType: "fifo",
			want:      cache.NewFIFOCache(),
			err:       nil,
		},
		{
			name:      "[Failure] Invalid Cache type",
			cacheType: "Random",
			want:      nil,
			err:       errors.New("Undefined cache type. Use one of the following: \"lru\",\"fifo\" or \"lfu\""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obtainedCache, err := GetCache(tt.cacheType)
			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, obtainedCache)
		})
	}

}
